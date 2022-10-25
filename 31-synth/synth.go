package main

// Hint: go run synth.go 220 440 880 512
// Hint: aplay -f S16_LE output.wav

import (
	"bufio"
	"encoding/binary"
	"flag"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

var flagSampleRate = flag.Int("r", 8000, "sample rate (in samples per second)")
var flagOutputFilename = flag.String("o", "output.wav", "wav file to write")

// See https://docs.fileformat.com/audio/wav/
type WavHeader struct {
	RiffMark           [4]byte
	Filesize           uint32
	WaveMark           [4]byte
	FmtChunk           [4]byte
	FmtChunkLen        uint32
	PcmType            uint16
	NumChans           uint16
	SampleRate         uint32
	Computed           uint32 // (Sample Rate * BitsPerSample * Channels) / 8
	DataBytesPerSample uint16
	BitsPerSample      uint16
	DataChunk          [4]byte
	DataChunkLen       uint32
}

// Mono, 16 bit, Little-Endian.
func WriteWavHeader(dataSize int, w io.Writer) {
	header := &WavHeader{
		RiffMark:           [4]byte{'R', 'I', 'F', 'F'},
		Filesize:           0,
		WaveMark:           [4]byte{'W', 'A', 'V', 'E'},
		FmtChunk:           [4]byte{'f', 'm', 't', ' '},
		FmtChunkLen:        16,
		PcmType:            1, // PCM=1
		NumChans:           1,
		SampleRate:         uint32(*flagSampleRate),
		Computed:           uint32(*flagSampleRate) * 16 * 1 / 8,
		DataBytesPerSample: 2,
		BitsPerSample:      16,
		DataChunk:          [4]byte{'d', 'a', 't', 'a'},
		DataChunkLen:       0,
	}
	err := binary.Write(w, binary.LittleEndian, header)
	if err != nil {
		log.Fatalf("cannot write wav header: %v", err)
	}
}

/* https://docs.fileformat.com/audio/wav/

1 - 4	“RIFF”	Marks the file as a riff file. Characters are each 1 byte long.
5 - 8	File size (integer)	Size of the overall file - 4 bytes, in bytes (32-bit integer). Typically, you’d fill this in after creation.
9 -12	“WAVE”	File Type Header. For our purposes, it always equals “WAVE”.
13-16	“fmt "	Format chunk marker. Includes trailing null
17-20	16	Length of format data as listed above
21-22	1	Type of format (1 is PCM) - 2 byte integer
23-24	2	Number of Channels - 2 byte integer
25-28	44100	Sample Rate - 32 byte integer. Common values are 44100 (CD), 48000 (DAT). Sample Rate = Number of Samples per second, or Hertz.
29-32	176400	(Sample Rate * BitsPerSample * Channels) / 8.
33-34	4	(BitsPerSample * Channels) / 8.1 - 8 bit mono2 - 8 bit stereo/16 bit mono4 - 16 bit stereo
35-36	16	Bits per sample
37-40	“data”	“data” chunk header. Marks the beginning of the data section.
41-44	File size (data)	Size of the data section.
*/

func main() {
	flag.Parse()

	f, err := os.Create(*flagOutputFilename)
	if err != nil {
		log.Fatalf("Cannot create %q: %v", *flagOutputFilename, err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	Synth(w)
}

func Synth(w io.Writer) {
	var voices []chan float64       // A holder for the voices.
	for _, a := range flag.Args() { // Create one voice per argument.
		voice := make(chan float64, 10000)
		go RunVoice(a, voice)
		voices = append(voices, voice)
	}

	mixed := make(chan float64, 10000)
	go RunMixer(voices, mixed) // The mixer adds the voice signals to make the mixed signal.

	for y := range mixed { // Convert each mixed sample to an int16 sample.
		sample := 10000 * y // Convert signal range [-1, 1] to [-10000, 10000].
		if sample > 32000 { // Clip if higher than this.
			sample = 32000
		}
		if sample < -32000 { // Clip if lower than this.
			sample = -32000
		}
		err := binary.Write(w, binary.LittleEndian, int16(sample))
		if err != nil {
			log.Fatalf("binary.Write fails: %v", err)
		}
	}
}

func RunVoice(a string, voice chan float64) {
	freq, err := strconv.ParseFloat(a, 64) // Parse the frequency.
	if err != nil {
		log.Fatalf("cannot ParseFloat %q: %v", a, err)
	}
	// Play for one second.
	for i := 0; i < *flagSampleRate; i++ {
		y := math.Sin(float64(i) * 2.0 * math.Pi * freq / float64(*flagSampleRate))
		voice <- y // Output the sine wave (range is -1 to 1).
	}
	close(voice)
}

func RunMixer(voices []chan float64, mixed chan float64) {
	for { // Until reading a voice fails because it is closed:
		sum := 0.0 // Sum each of the voices.
		for _, voice := range voices {
			y, ok := <-voice
			if !ok {
				close(mixed)
				return
			}
			sum += y
		}
		mixed <- sum // Write the sum to the mixed output.
	}
}
