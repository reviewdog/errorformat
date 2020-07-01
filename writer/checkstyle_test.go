package writer

import "os"

func ExampleCheckStyle() {
	w := NewCheckStyle(os.Stdout)
	for _, e := range testErrs {
		w.Write(e)
	}
	w.Flush()
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <checkstyle version="1.0">
	//   <file name="file2">
	//     <error column="14" line="2" message="emacs" severity="error 1" source="E1"></error>
	//     <error column="1" line="14" message="neovim" severity="error 14" source="E14"></error>
	//   </file>
	//   <file name="path/to/file1">
	//     <error column="14" line="1" message="hello" severity="warning"></error>
	//     <error column="14" line="2" message="vim" severity="info"></error>
	//   </file>
	// </checkstyle>
}
