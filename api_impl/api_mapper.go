package api_impl

import (
	"fmt"

	"github.com/lorenzo-milicia/go-server-queue/api"
	"github.com/lorenzo-milicia/go-server-queue/domain"
)

func toDto(input []domain.Record) []*api.Record {
	var records = make([]*api.Record, 0)

	for _, r := range input {
		fields := make([]*api.Field, 0)
		for k, v := range r.Fields {
			fields = append(fields,
				&api.Field{
					Name:  k,
					Value: fmt.Sprint(v),
				},
			)
		}

		records = append(records, &api.Record{Id: r.ID,Fields: fields})
	}

	return records
}
