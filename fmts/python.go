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

	register(&Fmt{
		Name: "mypy",
		Errorformat: []string{
			`%f:%l: %trror: %m`,
			`%f:%l: %tarning: %m`,
			`%f:%l: %tnfo: %m`,
			`%f:%l: %tote: %m`,
		},
		Description: "An optional static type checker for Python",
		URL:         "http://mypy-lang.org/",
		Language:    lang,
	})

	register(&Fmt{
		Name: "pydocstyle",
		Errorformat: []string{
			`%A%f:%l%r`,
			`%C%\s%+%m`,
		},
		Description: "A static analysis tool for checking compliance with Python docstring conventions",
		URL:         "https://github.com/PyCQA/pydocstyle",
		Language:    lang,
	})
	register(&Fmt{
		Name: "bandit",
		Errorformat: []string{
			`%f:%l: B%n[bandit]: %tIGH: %m`,
			`%f:%l: B%n[bandit]: %tEDIUM: %m`,
			`%f:%l: B%n[bandit]: %tOW: %m`,
		},
		Description: "A tool designed to find common security issues in Python code.",
		URL:         "https://github.com/PyCQA/bandit.git",
		Language:    lang,
	})
}
