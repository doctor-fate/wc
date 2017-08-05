package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// CountLines returns the number of newline characters, as well as standard wc utility does.
func CountLines(data []byte) int {
	return bytes.Count(data, []byte{'\n'})
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("filename not specified")
	}

	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("cannot read file %s: %s\n", filename, err)
	}

	fmt.Println(CountLines(data))
}
