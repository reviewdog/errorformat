package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/haya14busa/errorformat"
)

const usageMessage = "" +
	`Usage: errorformat [flags] [errorformat ...]

errorformat reads compiler/linter/static analyzer result from STDIN, formats
them by given 'errorformat' (90% compatible with Vim's errorformat. :h
errorformat), and outputs formated result to STDOUT.

Example:
	$ echo '/path/to/file:14:28: error message\nfile2:3:4: msg' | errorformat "%f:%l:%c: %m"
	/path/to/file|14 col 28| error message
	file2|3 col 4| msg
`

func usage() {
	fmt.Fprintln(os.Stderr, usageMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	errorformats := flag.Args()
	efm, err := errorformat.NewErrorformat(errorformats)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	s := efm.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(s.Entry())
	}
}
