package wave

import (
	"bufio"
	"fmt"
	"github.com/hendrixjl/gowave/utils"
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
	fs.subchunkID = utils.Read_fixed_string(r, SUBCHUNK_ID_SIZE)
	fs.size = utils.Read_uint32(r)
	fs.format = utils.Read_uint16(r)
	fs.numChannels = utils.Read_uint16(r)
	fs.sampleRate = utils.Read_uint32(r)
	fs.byteRate = utils.Read_uint32(r)
	fs.blockAlign = utils.Read_uint16(r)
	fs.bitsPerSample = utils.Read_uint16(r)

	return fs, FMT_SUBCHUNK_SIZE
}



