package wave

import (
	"bufio"
	"strings"
	"fmt"
)


type Subchunk interface {
    String() string
}


func Make_subchunk(r *bufio.Reader) (sc Subchunk, bytes_read int) {
	const ID_SIZE = 4
	peek_buf, err := r.Peek(ID_SIZE)
	if err != nil { panic(err) }
	if strings.Contains(string(peek_buf), "fmt ") {
		sc, bytes_read = Make_fmt_subchunk(r)
	} else if strings.Contains(string(peek_buf), "FLLR") {
		sc, bytes_read = Make_fllr_subchunk(r)
	} else if strings.Contains(string(peek_buf), "data") {
		sc, bytes_read = Make_data_subchunk(r)
	} else {
		fmt.Printf("Unknown ID = %s\n", string(peek_buf))
		bytes_read = 0
		sc = &Fllr_subchunk{ "Err ", 0 }
	}
	return sc, bytes_read
}

