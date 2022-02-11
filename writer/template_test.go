package writer

import (
	"os"
	"text/template"
)

func ExampleTemplate_string() {
	tmpl, _ := template.New("example").Parse("{{.String}}")
	w := NewTemplate(tmpl, os.Stdout)
	for _, e := range testErrs {
		w.Write(e)
	}
	// Output:
	// path/to/file1|1-2 col 14-15 warning| hello
	// path/to/file1|2 col 14 info| vim
	// file2|2 col 14 error 1| emacs
	// file2|14 col 1 error 14| neovim
}

func ExampleTemplate_more() {
	tmpl, _ := template.New("example").Parse("file:{{.Filename}}\tline:{{.Lnum}}\tcol:{{.Col}}\tmes:{{.Text}}")
	w := NewTemplate(tmpl, os.Stdout)
	for _, e := range testErrs {
		w.Write(e)
	}
	// Output:
	// file:path/to/file1	line:1	col:14	mes:hello
	// file:path/to/file1	line:2	col:14	mes:vim
	// file:file2	line:2	col:14	mes:emacs
	// file:file2	line:14	col:1	mes:neovim
}
