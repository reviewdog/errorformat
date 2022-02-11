package writer

import "os"

func ExampleJSONL() {
	w := NewJSONL(os.Stdout)
	for _, e := range testErrs {
		w.Write(e)
	}
	// Output:
	// {"filename":"path/to/file1","lnum":1,"end_lnum":2,"col":14,"end_col":15,"vcol":false,"nr":0,"pattern":"","text":"hello","type":87,"valid":false,"lines":["path/to/file1:1:14:[W] hello"]}
	// {"filename":"path/to/file1","lnum":2,"end_lnum":0,"col":14,"end_col":0,"vcol":false,"nr":0,"pattern":"","text":"vim","type":73,"valid":false,"lines":["path/to/file1:2:14:[I] vim"]}
	// {"filename":"file2","lnum":2,"end_lnum":0,"col":14,"end_col":0,"vcol":false,"nr":1,"pattern":"","text":"emacs","type":69,"valid":false,"lines":["file2:2:14:[E] emacs"]}
	// {"filename":"file2","lnum":14,"end_lnum":0,"col":1,"end_col":0,"vcol":false,"nr":14,"pattern":"","text":"neovim","type":69,"valid":false,"lines":["file2:14:1:[E] neovim"]}
}
