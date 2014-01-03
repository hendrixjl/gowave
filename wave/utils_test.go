
package wave

import (
	"testing" 
	"bufio"
	"fmt"
	"io"
)


type fakeReader struct {
	b bool
}

func (this *fakeReader) Read(p []byte) (n int, err error) {
	p[0] = 1
	p[1] = 2
	p[2] = 3
	p[3] = 4
	p[4] = 5
	p[5] = 6
	p[6] = 7
	p[7] = 8
	n = 8
	err = nil
	return
}

func TestRead_into_buffer_happy(t *testing.T) {
	var freader fakeReader
	reader := bufio.NewReader(&freader)
	buf := make([]byte, 5)
	bytesRead, err := Read_into_buffer(reader, 4, buf)
	if (err != nil) {
		t.Error(fmt.Sprintf("Error returned! (%s)", err));
	}
	if (bytesRead != 4) {
		t.Error(fmt.Sprintf("Wrong number of bytes Read! (%d)", bytesRead));
	}
	if buf[0] != 1 ||
	   buf[1] != 2 ||
	   buf[2] != 3 ||
	   buf[3] != 4 {
		t.Error("Values not correct!")
	}
}

func TestRead_into_buffer_tooSmall(t *testing.T) {
	var freader fakeReader
	reader := bufio.NewReader(&freader)
	buf := make([]byte, 2)
	bytesRead, err := Read_into_buffer(reader, 4, buf)
	if (err == nil) {
		t.Error("No error returned!")
	} else if (bytesRead != 0) {
		t.Error("No bytes should be read!")
	}
}

func TestRead_into_buffer_eofBeforeAllBytesRead(t *testing.T) {
	var freader fakeReader
	reader := bufio.NewReader(&freader)
	buf := make([]byte, 10)
	bytesRead, err := Read_into_buffer(reader, 10, buf)
	if (err == nil) {
		t.Error("No error was returned when one should have been!")
		return
	}
	if (err != io.EOF) {
		t.Error(fmt.Sprintf("Wrong error returned! (%v)", err))
	}
	if (bytesRead != 0) {
		t.Error("No bytes should be read!")
	}
}
