// +build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/haya14busa/errorformat/fmts"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	f, err := os.Create("doc.go")
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	fmt.Fprintln(w, pkgcomment)
	fmt.Fprintln(w, "//")
	fmt.Fprintln(w, "// Defined formats:")
	fmt.Fprintln(w, "// ")
	langToFmts := fmts.DefinedFmtsByLang()

	langs := make([]string, 0, len(langToFmts))
	for lang, _ := range langToFmts {
		langs = append(langs, lang)
	}
	sort.Strings(langs)

	for _, lang := range langs {
		nameToFmt := langToFmts[lang]
		names := make([]string, 0, len(nameToFmt))
		for name, _ := range nameToFmt {
			names = append(names, name)
		}
		sort.Strings(names)

		fmt.Fprintf(w, "// \t%v\n", lang)
		for _, name := range names {
			f := nameToFmt[name]
			fmt.Fprintf(w, "// \t\t%s\t%s - %s\n", f.Name, f.Description, f.URL)
		}
	}

	fmt.Fprintln(w, pkgline)
	return nil
}

const (
	pkgcomment = `// Package fmts holds defined errorformats.`
	pkgline    = `package fmts`
)
