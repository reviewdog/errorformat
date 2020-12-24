package fmts

func init() {
	const lang = "python"

	register(&Fmt{
		Name: "pep8",
		Errorformat: []string{
			`%f:%l:%c: %m`,
		},
		Description: "Python style guide checker",
		URL:         "https://pypi.python.org/pypi/pep8",
		Language:    lang,
	})

	register(&Fmt{
		Name: "flake8",
		Errorformat: []string{
			`%f:%l:%c: %t%n %m`,
		},
		Description: "Tool for python style guide enforcement",
		URL:         "https://flake8.pycqa.org/",
		Language:    lang,
	})

	register(&Fmt{
		Name: "black",
		Errorformat: []string{
			`%-GOh no!%.%#`,
			`%-G%\d\+ files%.%#`,
			`%-GAll done!%.%#`,
			`%m %f`,
			`%-G%.%#`,
		},
		Description: "A uncompromising Python code formatter",
		URL:         "https://github.com/psf/black",
		Language:    lang,
	})
}
