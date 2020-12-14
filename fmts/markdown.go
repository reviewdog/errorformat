package fmts

func init() {
	const lang = "markdown"

	register(&Fmt{
		Name: "remark-lint",
		Errorformat: []string{
			`%-P%f`,
			`%#%l:%c %# %trror  %m`,
			`%#%l:%c %# %tarning  %m`,
			`%-Q`,
			`%-G%.%#`,
		},
		Description: "Tool for writing clean and consistent markdown code",
		URL:         "https://github.com/remarkjs/remark-lint",
		Language:    lang,
	})
}
