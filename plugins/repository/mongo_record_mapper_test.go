package main

import (
	"testing"

	test "github.com/lorenzo-milicia/go-libs/testing"
	"github.com/lorenzo-milicia/go-server-queue/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestToDomain(t *testing.T) {
	var e = RecordEntity{
		ID:     primitive.NewObjectID(),
		Fields: map[string]interface{}{"field one": 1, "field two": "two"},
	}

	var d = e.toDomain()

	var expected = domain.Record{ID: e.ID.String(), Fields: e.Fields}

	test.AssertEquals(t, expected, d)
}

func BenchmarkToDomain(b *testing.B) {
	var e = RecordEntity{
		ID:     primitive.NewObjectID(),
		Fields: map[string]interface{}{"field one": 1, "field two": "two"},
	}

	var list = make([]domain.Record, 0)
	for n := 0; n < b.N; n++ {
		list = append(list, e.toDomain())
	}
}
