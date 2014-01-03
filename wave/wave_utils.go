package wave

import (
	"fmt"
	"bytes"
	"bufio"
	"io"
	"encoding/binary"
	"errors"
)

// read into a preallocated buffer
func Read_into_buffer(reader *bufio.Reader, size int, buf []byte) (totalRead int, err error) {
	if (len(buf) < size) {
		totalRead = 0
		err = errors.New("Input buffer is not large enough!")
	}
        totalRead = 0
	for totalRead < size {
		// read a chunk
		var n int
		n, err = reader.Read(buf[totalRead:size])
		if err != nil {
			return
		}
		totalRead += n
	}
	err = nil
	return
}

// read a certain number of bytes and return in a slice
func Read_bytes(reader *bufio.Reader, size int) []byte {
	buf := make([]byte, size)
	Read_into_buffer(reader, size, buf)
	return buf
}

func Read_fixed_string(reader *bufio.Reader, size int) string {
	return string(Read_bytes(reader, size))
}

// binary.Read only supports 64-bit integers
// adapt to 32 bit integer
func Read_uint32(reader *bufio.Reader) uint32 {
	const (
		UINT32_SIZE = 4
		UINT64_SIZE = 8
	)
	buf := make([]byte, UINT64_SIZE)
	Read_into_buffer(reader, UINT32_SIZE, buf)
	bufreader := bytes.NewReader(buf)
	var answer uint64
	err := binary.Read(bufreader, binary.LittleEndian, &answer)
	if err != nil {
		fmt.Println("binary.Read failed!")
		panic(err)
	}
	return uint32(answer)
}

// Adapt binary.Read to a 16 bit integer
func Read_uint16(reader *bufio.Reader) uint16 {
	const (
		UINT16_SIZE = 2
		UINT64_SIZE = 8
	)
	buf := make([]byte, UINT64_SIZE)
	Read_into_buffer(reader, UINT16_SIZE, buf)
	bufreader := bytes.NewReader(buf)
	var answer uint64
	err := binary.Read(bufreader, binary.LittleEndian, &answer)
	if err != nil {
		fmt.Println("binary.Read failed!")
		panic(err)
	}
	return uint16(answer)
}

