package fmts

func init() {
	const lang = "lua"

	register(&Fmt{
		Name: "luacheck",
		Errorformat: []string{
			`%f:%l:%c: %m`,
		},
		Description: "(luacheck --formatter=plain) Lua linter and static analyzer",
		URL:         "https://github.com/luarocks/luacheck",
		Language:    lang,
	})
}
