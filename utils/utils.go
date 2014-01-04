package utils

import (
	"fmt"
	"bytes"
	"bufio"
	"encoding/binary"
)

// read into a preallocated buffer
func Read_into_buffer(reader *bufio.Reader, buf []byte) (totalRead int, err error) {
	blen := len(buf)
	totalRead = 0
	for totalRead < blen {
		// read a chunk
		var n int
		n, err = reader.Read(buf[totalRead:blen])
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
	Read_into_buffer(reader, buf)
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
	sbuf := buf[0:UINT32_SIZE]
	Read_into_buffer(reader, sbuf)
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
	sbuf := buf[0:UINT16_SIZE]
	Read_into_buffer(reader, sbuf)
	bufreader := bytes.NewReader(buf)
	var answer uint64
	err := binary.Read(bufreader, binary.LittleEndian, &answer)
	if err != nil {
		fmt.Println("binary.Read failed!")
		panic(err)
	}
	return uint16(answer)
}
