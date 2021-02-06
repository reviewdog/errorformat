package fmts

func init() {
	const lang = "go"

	register(&Fmt{
		Name: "golint",
		Errorformat: []string{
			`%f:%l:%c: %m`,
		},
		Description: "linter for Go source code",
		URL:         "https://github.com/golang/lint",
		Language:    lang,
	})

	register(&Fmt{
		Name: "govet",
		Errorformat: []string{
			`%f:%l: %m`,
			`%-G%.%#`,
		},
		Description: "Vet examines Go source code and reports suspicious problems",
		URL:         "https://golang.org/cmd/vet/",
		Language:    lang,
	})

	register(&Fmt{
		Name: "golangci-lint",
		Errorformat: []string{
			`%E%f:%l:%c: %m`,
			`%E%f:%l: %m`,
			`%C%.%#`,
		},
		Description: "(golangci-lint run --out-format=line-number) GolangCI-Lint is a linters aggregator.",
		URL:         "https://github.com/golangci/golangci-lint",
		Language:    lang,
	})

	register(&Fmt{
		Name: "go-consistent",
		Errorformat: []string{
			`%f:%l:%c: %m`,
		},
		Description: "Source code analyzer that helps you to make your Go programs more consistent",
		URL:         "https://github.com/quasilyte/go-consistent",
		Language:    lang,
	})

	register(&Fmt{
		Name: "gosec",
		Errorformat: []string{
			`%f:%l:%c: %m`,
		},
		Description: "(gosec -fmt=golint) Golang Security Checker",
		URL:         "https://github.com/securego/gosec",
		Language:    lang,
	})

	register(&Fmt{
		Name:        "staticcheck",
		Errorformat: []string{
			"%f:%l:%c: %m",
		},
		Description: "Golang Static Analysis",
		URL:         "https://staticcheck.io",
		Language:    lang,
	})
}
