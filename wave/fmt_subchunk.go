package wave

import (
	"bufio"
	"fmt"
)

type Fmt_subchunk struct {
	subchunkID string
	size uint32
	format uint16
	numChannels uint16
	sampleRate uint32
	byteRate uint32
	blockAlign uint16
	bitsPerSample uint16
}

func (this *Fmt_subchunk) String() string {
	return fmt.Sprintf("[ID=%s, SIZE=%d, format=%d, numChannels=%d, sampleRate=%d, " +
			"byteRate=%d, blockAlign=%d, bitsPerSample=%d]", this.subchunkID, 
			this.size, this.format, this.numChannels, 
			this.sampleRate, this.byteRate, this.blockAlign, 
			this.bitsPerSample)
}



func Make_fmt_subchunk(r *bufio.Reader) (fs *Fmt_subchunk, bytes_read int) {
	const (
		SUBCHUNK_ID_SIZE = 4
		FMT_SUBCHUNK_SIZE = 24
	)

	fs = new(Fmt_subchunk)
	fs.subchunkID = Read_fixed_string(r, SUBCHUNK_ID_SIZE)
	fs.size = Read_uint32(r)
	fs.format = Read_uint16(r)
	fs.numChannels = Read_uint16(r)
	fs.sampleRate = Read_uint32(r)
	fs.byteRate = Read_uint32(r)
	fs.blockAlign = Read_uint16(r)
	fs.bitsPerSample = Read_uint16(r)

	return fs, FMT_SUBCHUNK_SIZE
}



