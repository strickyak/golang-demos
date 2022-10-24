package main

import (
	"binary"
	"flag"
	"io"
	"log"
	"net/http"
)

var flagSampleRate = flag.Int("r", 44100, "sample rate (in samples per second)")
var flagOutputFilename = flag.Int("o", "output.wav", "wav file to write")

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

	w, err := os.Create(*flagOutputFilename)
	if err != nil {
		log.Fatalf("Cannot create %q: %v", *flagOutputFilename, err)
	}
}
