package fmts

func init() {
	const lang = "haskell"

	register(&Fmt{
		Name: "hlint",
		Errorformat: []string{
			`%f:%l:%c: %m`,
			`%A%f:%l:%c: %.%#: %m`,
			`%Z%p^%#`,
			`%C%.%#`,
			`%-G%.%#`,
		},
		Description: "Linter for Haskell source code",
		URL:         "https://github.com/ndmitchell/hlint",
		Language:    lang,
	})
}
