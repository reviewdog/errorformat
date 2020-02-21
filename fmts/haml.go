package fmts

func init() {
	const lang = "haml"

	register(&Fmt{
		Name: "haml-lint",
		Errorformat: []string{
			`%f:%l [%t] %m`,
			`%-G%.%#`,
		},
		Description: "Tool for writing clean and consistent HAML",
		URL:         "https://github.com/sds/haml-lint",
		Language:    lang,
	})
}
