package fmts

func init() {
	const lang = "haskell"

	register(&Fmt{
		Name: "hlint",
		Errorformat: []string{
			`%f:(%l,%c)-(%e,%k): %tarning: %A%m`,
			`%f:(%l,%c)-(%e,%k): %trror: %A%m`,
			`%f:(%l,%c)-(%e,%k): %tuggestion: %A%m`,
			`%f:%l:%c-%k: %tarning: %A%m`,
			`%f:%l:%c-%k: %trror: %A%m`,
			`%f:%l:%c-%k: %tuggestion: %A%m`,
			`%f:%l:%c: %tarning: %A%m`,
			`%f:%l:%c: %trror: %A%m`,
			`%f:%l:%c: %tuggestion: %A%m`,
			`%C%m`,
			`%-G%.%#`
		},
		Description: "Linter for Haskell source code",
		URL:         "https://github.com/ndmitchell/hlint",
		Language:    lang,
	})
}
