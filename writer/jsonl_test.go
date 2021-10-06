package writer

import "os"

func ExampleJSONL() {
	w := NewJSONL(os.Stdout)
	for _, e := range testErrs {
		w.Write(e)
	}
	// Output:
	// {"filename":"path/to/file1","lnum":1,"col":14,"vcol":false,"nr":0,"pattern":"","text":"hello","type":"W","valid":false,"lines":["path/to/file1:1:14:[W] hello"]}
	// {"filename":"path/to/file1","lnum":2,"col":14,"vcol":false,"nr":0,"pattern":"","text":"vim","type":"I","valid":false,"lines":["path/to/file1:2:14:[I] vim"]}
	// {"filename":"file2","lnum":2,"col":14,"vcol":false,"nr":1,"pattern":"","text":"emacs","type":"E","valid":false,"lines":["file2:2:14:[E] emacs"]}
	// {"filename":"file2","lnum":14,"col":1,"vcol":false,"nr":14,"pattern":"","text":"neovim","type":"E","valid":false,"lines":["file2:14:1:[E] neovim"]}
}
