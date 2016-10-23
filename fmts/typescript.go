package fmts

func init() {
	const lang = "typescript"

	register(&Fmt{
		Name: "tslint",
		Errorformat: []string{
			`%f[%l, %c]: %m`, // --format=prose
		},
		Description: "An extensible linter for the TypeScript language",
		URL:         "https://github.com/palantir/tslint",
		Language:    lang,
	})

}
