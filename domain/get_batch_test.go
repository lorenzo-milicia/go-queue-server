package domain

import (
	"testing"

	test "github.com/lorenzo-milicia/go-libs/testing"
)

var mockRecords = []Record{
	{
		ID: "id",
		Fields: map[string]interface{}{
			"one": 1,
		},
	},
}

func TestGetBatch(t *testing.T) {
	var pagesize int = 100

	var repo MockRepository = MockRepository{}

	var service IBatchService = BatchService{Repository: &repo}

	var result = service.GetBatch(pagesize, 0)

	if repo.findAllPaginatedCount != 1 {
		t.Error("Expected one call to the method FindAllPaginated")
	}

	test.AssertEquals(t, mockRecords, result)
}

type MockRepository struct {
	findAllPaginatedCount int
}

func (r *MockRepository) FindAllPaginated(pagesize int, pagenumber int) []Record {
	r.findAllPaginatedCount++
	return mockRecords
}
