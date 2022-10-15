package domain

type Record struct {
	ID     string
	Fields map[string]interface{}
}

type Records []Record