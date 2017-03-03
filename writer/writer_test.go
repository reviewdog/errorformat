package writer

import "github.com/haya14busa/errorformat"

var errors = []*errorformat.Entry{
	{Filename: "path/to/file1", Lnum: 1, Col: 14, Text: "hello", Type: 'W'},
	{Filename: "path/to/file1", Lnum: 2, Col: 14, Text: "vim", Type: 'I'},
	{Filename: "file2", Lnum: 2, Col: 14, Text: "emacs", Type: 'E', Nr: 1},
	{Filename: "file2", Lnum: 14, Col: 1, Text: "neovim", Type: 'E', Nr: 14},
}
