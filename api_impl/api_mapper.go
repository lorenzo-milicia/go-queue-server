package api_impl

import (
	"fmt"

	"github.com/lorenzo-milicia/go-server-queue/api"
	"github.com/lorenzo-milicia/go-server-queue/domain"
)

func toDto(r domain.Record) *api.Record {
	fields := make([]*api.Field, 0)
	for k, v := range r.Fields {
		fields = append(fields,
			&api.Field{
				Name:  k,
				Value: fmt.Sprint(v),
			},
		)
	}

	return &api.Record{Id: r.ID, Fields: fields}
}

func toDomain(r *api.Record) *domain.Record {
	return &domain.Record{ID: r.Id, Payload: r.Payload}
}