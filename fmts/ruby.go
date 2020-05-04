package fmts

func init() {
	const lang = "ruby"

	register(&Fmt{
		Name: "brakeman",
		Errorformat: []string{
			`%f%*\s%l%*\s%m`,
		},
		Description: "(brakeman --quiet --format tabs) A static analysis security vulnerability scanner for Ruby on Rails applications",
		URL:         "https://github.com/presidentbeef/brakeman",
		Language:    lang,
	})

	register(&Fmt{
		Name: "fasterer",
		Errorformat: []string{
			`%f:%l %m`,
			`%-G%.%#`,
		},
		Description: "Speed improvements suggester",
		URL:         "https://github.com/DamirSvrtan/fasterer",
		Language:    lang,
	})

	register(&Fmt{
		Name: "reek",
		Errorformat: []string{
			`%*\s%f:%l: %m`,
			`%-G%.%#`,
		},
		Description: "(reek --single-line) Code smell detector for Ruby",
		URL:         "https://github.com/troessner/reek",
		Language:    lang,
	})

	register(&Fmt{
		Name: "rubocop",
		Errorformat: []string{
			`%A%f:%l:%c: %t: %m`,
			`%Z%p^%#`,
			`%C%.%#`,
			`%-G%.%#`,
		},
		Description: "A Ruby static code analyzer, based on the community Ruby style guide",
		URL:         "https://github.com/rubocop-hq/rubocop",
		Language:    lang,
	})

	register(&Fmt{
		Name: "sorbet",
		Errorformat:[]string{
			`%f:%l: %m`,
			`%-G%.%#`,
		},
		Description: "A fast, powerful type checker designed for Ruby",
		URL:         "https://github.com/sorbet/sorbet",
		Language:    lang,
	})

	register(&Fmt{
		Name: "standardrb",
		Errorformat: []string{
			`%*\s%f:%l:%c: %.%#: %m`,
			`%-G%.%#`,
		},
		Description: "(standard) Ruby style guide, linter, and formatter",
		URL:         "https://github.com/testdouble/standard",
		Language:    lang,
	})
}
