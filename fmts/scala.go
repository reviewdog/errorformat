package fmts

func init() {
	const lang = "scala"

	register(&Fmt{
		Name: "scalac",
		Errorformat: []string{
			`%E%f:%l: error: %m`,
			`%W%f:%l: warning: %m`,
			`%A%f:%l: %m`,
			`%Z%p^`,
			`%C%.%#`,
			`%-G%.%#`,
		},
		Description: "Scala compiler",
		URL:         "http://www.scala-lang.org/",
		Language:    lang,
	})
}
