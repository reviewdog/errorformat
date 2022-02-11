package writer

import (
	"fmt"

	"github.com/reviewdog/errorformat"
)

var testErrs = []*errorformat.Entry{
	{Filename: "path/to/file1", Lnum: 1, EndLnum: 2, Col: 14, EndCol: 15, Text: "hello", Type: 'W'},
	{Filename: "path/to/file1", Lnum: 2, Col: 14, Text: "vim", Type: 'I'},
	{Filename: "file2", Lnum: 2, Col: 14, Text: "emacs", Type: 'E', Nr: 1},
	{Filename: "file2", Lnum: 14, Col: 1, Text: "neovim", Type: 'E', Nr: 14},
}

func init() {
	for _, e := range testErrs {
		e.Lines = append(e.Lines, fmt.Sprintf("%s:%d:%d:[%s] %s",
			e.Filename, e.Lnum, e.Col, string(e.Type), e.Text))
	}
}
