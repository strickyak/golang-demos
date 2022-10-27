package main

// Hint: go run synth.go 220 440 880 512
// Hint: aplay -f S16_LE output.wav

// Octave Naming:
// https://www.allaboutmusictheory.com/piano-keyboard/music-note-names/

// A4 is 440 Hz:
// https://en.wikipedia.org/wiki/A440_(pitch_standard)

// Ode to Joy
// https://www.ibiblio.org/mutopia/ftp/BeethovenLv/ode/ode-let.pdf

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

var flagSampleRate = flag.Int("r", 8000, "sample rate (in samples per second)")
var flagTranspose = flag.Float64("t", 0, "transpose this many steps up or down")
var flagOutputFilename = flag.String("o", "output.wav", "wav file to write")
var flagNoWavHeader = flag.Bool("n", false, "omit WAV header; just raw S16_LE")

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatalf("At least 1 command line argument is required.")
	}

	file, err := os.Create(*flagOutputFilename)
	if err != nil {
		log.Fatalf("Cannot create %q: %v", *flagOutputFilename, err)
	}
	defer file.Close()

	InitTones()
	w := bufio.NewWriter(file)
	if !*flagNoWavHeader {
		WriteWavHeader(900*(*flagSampleRate), w) // Lie and say 900 seconds (15 min).
	}
	Synth(w)
	w.Flush()
}

func Synth(w io.Writer) {
	var voices []chan float64         // A holder for the voices.
	for _, arg := range flag.Args() { // Create one voice per argument.
		voice := make(chan float64, 10000)
		go RunVoice(arg, voice)
		voices = append(voices, voice)
	}

	mixed := make(chan float64, 10000)
	go RunMixer(voices, mixed) // The mixer adds the voice signals to make the mixed signal.

	for y := range mixed { // Convert each mixed sample to an int16 sample.
		sample := 8000 * y  // Convert signal float [-1.0, 1.0] to uint16 [-8000, 8000].
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

func RunVoice(arg string, voice chan float64) {
	for _, a := range strings.Split(arg, ",") {
		freq, ok := tones[a]
		if !ok {
			log.Fatalf("Unknown tone %q", a)
		}
		// Play for half a second.
		for i := 0; i < *flagSampleRate/2; i++ {
			y := math.Sin(float64(i) * 2.0 * math.Pi * freq / float64(*flagSampleRate))
			voice <- y // Output the sine wave (range is -1 to 1).
		}
		// One tenth of a second gap.
		for i := 0; i < *flagSampleRate/10; i++ {
			voice <- 0
		}
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

////  Begin Tone Math

var (
	chromaticScale = []string{
		"a", "a#", "b", "c", "c#", "d",
		"d#", "e", "f", "f#", "g", "g#",
	}
	tones = make(map[string]float64)
)

func InitTones() {
	octave := 0
	halfStep := *flagTranspose - 4*12 // 4 octaves below the anchor A4 == 440 Hz
	for octave < 9 {
		for _, note := range chromaticScale {
			if note == "c" {
				octave++
			}
			noteAndOctave := fmt.Sprintf("%s%d", note, octave) // e.g. c4
			// How the even-tempered scale works:
			tones[noteAndOctave] = 440.0 * math.Pow(2.0, float64(halfStep)/12.0)
			halfStep++
		}
	}
	tones["_"] = 0.0 // Add "_" for a rest, with frequency 0.
	// log.Printf("Tones: %#v", tones)
}

////////////// WAV Header

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
