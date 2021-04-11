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
			`%-G%\d\+ file%.%#`,
			`%-GAll done!%.%#`,
			`%trror:%f:%l:%c:%m`,
			`%trror:%m`,
			`%m %f`,
			`%-G%.%#`,
		},
		Description: "A uncompromising Python code formatter",
		URL:         "https://github.com/psf/black",
		Language:    lang,
	})

	register(&Fmt{
		Name: "isort",
		Errorformat: []string{
			`%tRROR: %f %m`,
			`%-GSkipped %\d\+ file%.`,
		},
		Description: "A Python utility / library to sort Python imports",
		URL:         "https://github.com/PyCQA/isort",
		Language:    lang,
	})
}
