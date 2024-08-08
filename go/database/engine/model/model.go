package model

type Data struct {
	Engine        string
	ConnStr       string
	Queries       string
    BorderStyle   int
	DestFolder    string
	LuaTabulaPath string
	TabulaLogFile string
	Option        Option
}

type Option int

const (
	RUN Option = iota + 1
	TABLES
)
