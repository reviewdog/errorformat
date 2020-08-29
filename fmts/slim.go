package fmts

func init() {
	const lang = "slim"

	register(&Fmt{
		Name: "slim-lint",
		Errorformat: []string{
			`%f:%l [%t] %m`,
			`%-G%.%#`,
		},
		Description: "Tool to help keep your Slim files clean and readable",
		URL:         "https://github.com/sds/slim-lint",
		Language:    lang,
	})
}
