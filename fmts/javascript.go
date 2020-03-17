package fmts

func init() {
	const lang = "javascript"

	register(&Fmt{
		Name: "eslint",
		Errorformat: []string{
			`%-P%f`,
			` %#%l:%c %# %trror  %m`,
			` %#%l:%c %# %tarning  %m`,
			`%-Q`,
			`%-G%.%#`,
		},
		Description: "(eslint [-f stylish]) A fully pluggable tool for identifying and reporting on patterns in JavaScript",
		URL:         "https://github.com/eslint/eslint",
		Language:    lang,
	})

	register(&Fmt{
		Name: "eslint-compact",
		Errorformat: []string{
			`%f: line %l, col %c, %trror - %m`,
			`%f: line %l, col %c, %tarning - %m`,
			`%-G%.%#`,
		},
		Description: "(eslint -f compact) A fully pluggable tool for identifying and reporting on patterns in JavaScript",
		URL:         "https://github.com/eslint/eslint",
		Language:    lang,
	})

	register(&Fmt{
		Name: "standardjs",
		Errorformat: []string{
			`%*\s%f:%l:%c: %m`,
			`%-G%.%#`,
		},
		Description: "(standard) JavaScript style guide, linter, and formatter",
		URL:         "https://github.com/standard/standard",
		Language:    lang,
	})

}
