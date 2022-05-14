package main

import (
	"fmt"

	"github.com/lorenzo-milicia/go-server-queue/domain"
)

func (e RecordEntity) toDomain() domain.Record {
	return domain.Record{
		ID:     fmt.Sprint(e.ID.Hex()),
		Fields: e.Fields,
	}
}
