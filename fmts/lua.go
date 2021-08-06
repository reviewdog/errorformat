package fmts

func init() {
	const lang = "lua"

	register(&Fmt{
		Name: "luacheck",
		Errorformat: []string{
			`%f:%l:%c: %m`,
		},
		Description: "a linter and a static analyzer for Lua",
		URL:         "https://github.com/luarocks/luacheck",
		Language:    lang,
	})
}
