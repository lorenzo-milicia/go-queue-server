package repository

import (
	"fmt"

	"github.com/lorenzo-milicia/go-server-queue/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecordEntity struct {
	ID     primitive.ObjectID     `bson:"_id"`
	Fields map[string]interface{} `bson:"data"`
}

func (e RecordEntity) toDomain() domain.Record {
	return domain.Record{
		ID:     fmt.Sprint(e.ID.Hex()),
		Fields: e.Fields,
	}
}
