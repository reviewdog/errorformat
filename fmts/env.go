package fmts

func init() {
	const lang = "env"

	register(&Fmt{
		Name: "dotenv-linter",
		Errorformat: []string{
			`%f:%l %m`,
		},
		Description: "Lightning-fast linter for .env files. Written in Rust",
		URL:         "https://github.com/dotenv-linter/dotenv-linter",
		Language:    lang,
	})
}
