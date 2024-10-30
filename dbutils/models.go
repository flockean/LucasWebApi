package dbutils

type Project struct {
	id          int8
	name        string
	description string
}

type Service struct {
	id      int8
	name    string
	lang    string
	focus   string
	project int
}
