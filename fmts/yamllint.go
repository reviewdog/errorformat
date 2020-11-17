package fmts

func init() {
	const lang = "yaml"

	register(&Fmt{
		Name: "yamllint",
		Errorformat: []string{
			`%f:%l:%c: %m`,
		},
		Description: "linter for YAML files",
		URL:         "https://github.com/adrienverge/yamllint",
		Language:    lang,
	})
}
