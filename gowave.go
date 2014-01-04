
package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/hendrixjl/gowave/wave"
)

func main() {
	args := os.Args
	if (len(args) != 2) {
		fmt.Printf("Usage: %s <filename>\n", args[0])
		return
	}
	fmt.Printf("Argument is %s\n", args[1])

	// open input file
	fi, err := os.Open(args[1])
	if err != nil { panic(err) }
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// make a read buffer
	r := bufio.NewReader(fi)

	whead, _ := wave.Make_riff_header(r)

	fmt.Printf("header = %v\n", whead)

	sizeLeft := whead.GetDataSize() - 4

	for sizeLeft > 0 {
		subchunk, bytes_read := wave.Make_subchunk(r)
		if (bytes_read > 0) {
			fmt.Println(subchunk);
			sizeLeft -= bytes_read;
		} else {
			fmt.Print("Error in reading!");
			break;
		}
	}
}

