package fmts

func init() {
	const lang = "protocolbuffer"

	register(&Fmt{
		Name: "protolint",
		Errorformat: []string{
			`[%f:%l:%c] %m`,
		},
		Description: "A pluggable linting utility for Protocol Buffer files",
		URL:         "https://github.com/yoheimuta/protolint",
		Language:    lang,
	})
}
