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

	register(&Fmt{
		Name: "buf",
		Errorformat: []string{
			`%f:%l:%c:%m`,
		},
		Description: "A new way of working with Protocol Buffers.",
		URL:         "https://github.com/bufbuild/buf",
		Language:    lang,
	})
}
