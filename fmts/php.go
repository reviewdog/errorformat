package fmts

func init() {
	const lang = "php"

	register(&Fmt{
		Name: "phpstan",
		Errorformat: []string{
			`%f:%l:%m`,
		},
		Description: "(phpstan --error-format=raw) PHP Static Analysis Tool - discover bugs in your code without running it!",
		URL:         "https://github.com/phpstan/phpstan",
		Language:    lang,
	})

	register(&Fmt{
		Name: "psalm",
		Errorformat: []string{
			`%f:%l:%c:%m`,
		},
		Description: "(psalm --output-format=text) Psalm is a static analysis tool for finding errors in PHP",
		URL:         "https://github.com/vimeo/psalm",
		Language:    lang,
	})
}
