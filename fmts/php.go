package fmts

func init() {
	const lang = "php"

	register(&Fmt{
		Name: "phpstan",
		Errorformat: []string{
			`%f:%l:%m`,
		},
		Description: "(phpstan --errorFormat=raw) PHP Static Analysis Tool - discover bugs in your code without running it!",
		URL:         "https://github.com/phpstan/phpstan",
		Language:    lang,
	})

	register(&Fmt{
		Name: "phpcs",
		Errorformat: []string{
			`%f:%l:%c: %s - %m`,
		},
		Description: "(phpcs --report=emacs -q) PHP_CodeSniffer - tokenizes PHP files to detect violations of a defined coding standard!",
		URL:         "https://github.com/squizlabs/PHP_CodeSniffer",
		Language:    lang,
	})
}
