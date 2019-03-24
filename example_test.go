package errorformat_test

import (
	"fmt"
	"strings"

	"github.com/reviewdog/errorformat"
)

func ExampleErrorformat() {
	in := `
golint.new.go:3:5: exported var V should have comment or be unexported
golint.new.go:5:5: exported var NewError1 should have comment or be unexported
golint.new.go:7:1: comment on exported function F should be of the form "F ..."
golint.new.go:11:1: comment on exported function F2 should be of the form "F2 ..."
`
	efm, _ := errorformat.NewErrorformat([]string{`%f:%l:%c: %m`, `%-G%.%#`})
	s := efm.NewScanner(strings.NewReader(in))
	for s.Scan() {
		fmt.Println(s.Entry())
	}
	// Output:
	// golint.new.go|3 col 5| exported var V should have comment or be unexported
	// golint.new.go|5 col 5| exported var NewError1 should have comment or be unexported
	// golint.new.go|7 col 1| comment on exported function F should be of the form "F ..."
	// golint.new.go|11 col 1| comment on exported function F2 should be of the form "F2 ..."
}
