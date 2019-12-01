package fmts

func init() {
	const lang = "rust"

	register(&Fmt{
		Name: "clippy",
		Errorformat: []string{
			`%f:%l:%c: %m`,
		},
		Description: "(cargo clippy -q --message-format=short) A bunch of lints to catch common mistakes and improve your Rust code",
		URL:         "https://github.com/rust-lang/rust-clippy",
		Language:    lang,
	})
}
