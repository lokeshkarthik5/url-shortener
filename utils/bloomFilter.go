package utils

import (
	"context"

	"github.com/bits-and-blooms/bloom/v3"
	"github.com/lokeshkarthik5/url-shortner/internal/database"
)

type Bloom struct {
	userBloom *bloom.BloomFilter
	db        *database.Queries
	context   context.Context
}

var linkBloom *bloom.BloomFilter

func InitBloom() {
	linkBloom = bloom.NewWithEstimates(5_000_000, 0.001)
}

func PopulateBloom(ctx context.Context, q *database.Queries) error {
	shorts, err := q.PopulateBloom(ctx)
	if err != nil {
		return err
	}
	for _, short := range shorts {
		linkBloom.AddString(short)
	}

	return nil

}
