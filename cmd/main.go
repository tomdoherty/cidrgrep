package main

import (
	"log"
	"os"

	"github.com/tomdoherty/cidrgrep"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("%s <network/cidr> [files..]", os.Args[0])
	}

	var prefix string

	// process files passed on the command line
	if len(os.Args) > 2 {
		for _, file := range os.Args[2:] {
			if len(os.Args) > 3 {
				// when parsing multiple files, prefix with the filename
				prefix = file + ": "
			}
			f, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			cidrgrep.Filter(f, os.Stdout, os.Args[1], prefix)
		}
	} else {
		// default to reading stdin
		cidrgrep.Filter(os.Stdin, os.Stdout, os.Args[1], prefix)
	}
}
