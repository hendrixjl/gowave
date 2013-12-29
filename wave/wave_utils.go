package wave

import (
	"fmt"
	"bytes"
	"bufio"
	"io"
	"encoding/binary"
)

// read into a preallocated buffer
func Read_into_buffer(reader *bufio.Reader, size int, buf []byte) {
	for totalRead := 0; totalRead < size; {
		// read a chunk
		n, err := reader.Read(buf[totalRead:size])
		if err != nil && err != io.EOF { panic(err) }
		totalRead += n
	}
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

