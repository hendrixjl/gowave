package wave

import (
	"bufio"
	"fmt"
	"github.com/hendrixjl/gowave/utils"
)

type Riff_header struct {
	chunkID string
	chunkSize uint32
	format string
}

func (this *Riff_header) String() string {
	return fmt.Sprintf("[ID=%s, SIZE=%d, FORMAT=%s]", this.chunkID, this.chunkSize, this.format)
}

func (this *Riff_header) GetDataSize() int {
	return int(this.chunkSize)
}

func Make_riff_header(r *bufio.Reader) (rh *Riff_header, bytes_read int) {
	const (
		CHUNK_ID_SIZE = 4
		FORMAT_SIZE = 4
		RIFF_HEADER_SIZE = 4
	)

	rh = new(Riff_header)	
	rh.chunkID = utils.Read_fixed_string(r, CHUNK_ID_SIZE)
	rh.chunkSize = utils.Read_uint32(r)
	rh.format = utils.Read_fixed_string(r, FORMAT_SIZE)

	bytes_read = RIFF_HEADER_SIZE
	return rh, bytes_read
}
