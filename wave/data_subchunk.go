package wave

import (
	"bufio"
	"fmt"
	"github.com/hendrixjl/gowave/utils"
)

type Data_subchunk struct {
	subchunkID string
	size uint32
}


func (this *Data_subchunk) String() string {
	return fmt.Sprintf("[ID=%s, SIZE=%d]", this.subchunkID, this.size)
}


func Make_data_subchunk(r *bufio.Reader) (fs *Data_subchunk, bytes_read int) {
	const (
		SUBCHUNK_ID_SIZE = 4
	)

	fs = new(Data_subchunk)
	fs.subchunkID = utils.Read_fixed_string(r, SUBCHUNK_ID_SIZE)
	fs.size = utils.Read_uint32(r)

        _ = utils.Read_bytes(r, int(fs.size))
	bytes_read = int(fs.size) + 8

	return fs, bytes_read
}



