package api_impl

import (
	"testing"

	"github.com/lorenzo-milicia/go-server-queue/domain"
)

func BenchmarkToDto(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := make([]domain.Record, n)
		toDto(r)
	}
}
