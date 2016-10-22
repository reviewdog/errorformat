package fmts

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/haya14busa/errorformat"
)

func TestFmts(t *testing.T) {
	for name, f := range DefinedFmts() {
		testfmt(t, name, f)
	}
}

func testfmt(t *testing.T, name string, f *Fmt) {
	t.Log(name)
	infile := fmt.Sprintf("testdata/%s.in", name)
	in, err := os.Open(infile)
	if err != nil {
		t.Error("no test for %q: %v", name, err)
		return
	}
	defer in.Close()
	okfile := fmt.Sprintf("testdata/%s.ok", name)
	ok, err := os.Open(okfile)
	if err != nil {
		t.Error("no ok test for %q: %v", name, err)
		return
	}
	defer ok.Close()
	outfile := fmt.Sprintf("testdata/%s.out", name)
	out, err := os.Create(outfile)
	if err != nil {
		t.Error(err)
		return
	}
	defer out.Close()
	efm, err := errorformat.NewErrorformat(f.Errorformat)
	if err != nil {
		t.Error(err)
	}
	bufout := bufio.NewWriter(out)
	s := efm.NewScanner(in)
	for s.Scan() {
		bufout.WriteString(s.Entry().String() + "\n")
	}
	if err := bufout.Flush(); err != nil {
		t.Error(err)
	}
	b, err := exec.Command("diff", "-u", okfile, outfile).Output()
	if err != nil {
		t.Error(err)
	}
	if d := string(b); d != "" {
		t.Error(d)
	}
}
