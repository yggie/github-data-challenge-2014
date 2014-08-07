package models

type _Language string
type LanguageDistribution map[_Language]float64
type Languages map[_Language]int64

const (
	C          = _Language("C")
	CPP        = _Language("C++")
	JAVA       = _Language("Java")
	JAVASCRIPT = _Language("Javascript")
	SCALA      = _Language("Scala")
	RUBY       = _Language("Ruby")
	PYTHON     = _Language("Python")
	GO         = _Language("Go")
	R          = _Language("R")
	SHELL      = _Language("Shell")
)
