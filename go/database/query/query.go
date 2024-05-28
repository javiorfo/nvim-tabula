package query

type ColumnResult struct {
	Name   string
	Length int
}

type SelectResult struct {
	Header map[int]ColumnResult
	Rows   map[int][]string
}
