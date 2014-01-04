
package utils

import (
	"testing" 
	"bufio"
	"fmt"
	"io"
)


type fakeReader struct {
	size int
}

func (this *fakeReader) Read(p []byte) (n int, err error) {
	if (this.size <= 0) {
		n = 0
		err = io.EOF
	} else if len(p) >= this.size {
		n = this.size
		for i:=0; i<n; i++ {
			p[i] = byte(i+this.size)
		}
		this.size = 0
		err = nil
	} else {
		n = len(p)
		for i:=0; i<n; i++ {
			p[i] = byte(i+this.size)
		}
		this.size -= n
		err = nil
	}
	return
}

func TestRead_into_buffer_happy(t *testing.T) {
	freader := fakeReader{10}
	reader := bufio.NewReader(&freader)
	buf := make([]byte, 5)
	bytesRead, err := Read_into_buffer(reader, 4, buf)
	if (err != nil) {
		t.Error(fmt.Sprintf("Error returned! (%s)", err));
	}
	if (bytesRead != 4) {
		t.Error(fmt.Sprintf("Wrong number of bytes Read! (%d)", bytesRead));
	}
	if buf[0] != 10 ||
	   buf[1] != 11 ||
	   buf[2] != 12 ||
	   buf[3] != 13 {
		t.Error("Values not correct!")
	}
}

func TestRead_into_buffer_tooSmall(t *testing.T) {
	freader := fakeReader{10}
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
	freader := fakeReader{5}
	reader := bufio.NewReader(&freader)
	buf := make([]byte, 10)
	bytesRead, err := Read_into_buffer(reader, 10, buf)
	if (err == nil) {
		t.Error(fmt.Sprintf("No error was returned when one should have been! bytesRead=%d", bytesRead))
		return
	}
	if (err != io.EOF) {
		t.Error(fmt.Sprintf("Wrong error returned! (%v)", err))
	}
	if (bytesRead != 5) {
		t.Error(fmt.Sprintf("Only 5 bytes should have been read. Instead, read %d bytes!", bytesRead))
	}
}


func TestRead_bytes_happy(t *testing.T) {
	freader := fakeReader{10}
	reader := bufio.NewReader(&freader)
	buf := Read_bytes(reader, 4)
	if len(buf) != 4 {
		t.Error(fmt.Sprintf("Wrong number of bytes Read! (%d)", len(buf)));
	}
	if buf[0] != 10 ||
	   buf[1] != 11 ||
	   buf[2] != 12 ||
	   buf[3] != 13 {
		t.Error("Values not correct!")
	}
}


type fakeStringReader struct {
	t string
}

func (this *fakeStringReader) Read(p []byte) (n int, err error) {
	byteArray := []byte(this.t)
	count := len(byteArray)

	if (count <= 0) {
		n = 0
		err = io.EOF
	} else if len(p) >= count {
		n = count
		for i:=0; i<n; i++ {
			p[i] = byteArray[i]
		}
		this.t = ""
		err = nil
	} else {
		n = len(p)
		for i:=0; i<n; i++ {
			p[i] = byteArray[i]
		}
// @TODO chop off the first n bytes of 
		err = nil
	}
	return
}

func TestRead_fixed_string_happy(t *testing.T) {
	freader := fakeStringReader{"How are you"}
	reader := bufio.NewReader(&freader)
	s := Read_fixed_string(reader, 4)
	if s != "How " {
		t.Error("String value is not correct! (%s)", s)
	}
}




