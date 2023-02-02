package fmts

func init() {
	const lang = "common"

	register(&Fmt{
		Name: "misspell",
		Errorformat: []string{
			`%f:%l:%c: %m`,
		},
		Description: "Correct commonly misspelled English words in source files",
		URL:         "https://github.com/client9/misspell",
		Language:    lang,
	})

	register(&Fmt{
		Name: "typos",
		Errorformat: []string{
			`%Eerror: %m`,
			`%C  --> %f:%l:%c%Z`,
			`%-G%.%#`,
		},
		Description: "Source code spell checker",
		URL:         "https://github.com/crate-ci/typos",
		Language:    lang,
	})
}
