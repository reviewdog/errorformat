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

	register(&Fmt{
		Name: "sbt",
		Errorformat: []string{
			`%E[%t%.%+] %f:%l: error: %m`,
			`%A[%t%.%+] %f:%l: %m`,
			`%Z[%.%+] %p^`,
			`%C[%.%+] %.%#`,
			`%-G%.%#`,
		},
		Description: "the interactive build tool",
		URL:         "http://www.scala-sbt.org/",
		Language:    lang,
	})
}
