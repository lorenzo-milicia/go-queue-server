package domain

type Record struct {
	ID     string
	Fields map[string]interface{}
	Payload string
}

type Records []*Record