package fmts

func init() {
	const lang = "csharp"

	register(&Fmt{
		Name: "msbuild",
		Errorformat: []string{
			`\ %#%f(%l\\\,%c):\ %m`,
		},
		Description: "(msbuild /property:GenerateFullPaths=true /nologo /v:q) Microsoft Build Engine",
		URL:         "https://docs.microsoft.com/en-us/visualstudio/msbuild/msbuild",
		Language:    lang,
	})
	register(&Fmt{
		Name: "dotnet",
		Errorformat: []string{
			`\ %#%f(%l\\\,%c):\ %m`,
		},
		Description: "(dotnet build -clp:NoSummary -p:GenerateFullPaths=true --no-incremental --nologo -v q) .NET Core CLI",
		URL:         "https://docs.microsoft.com/en-us/dotnet/core/tools/",
		Language:    lang,
	})
}
