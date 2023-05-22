package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/graph-gophers/dataloader"
)

var exampleBatchFn dataloader.BatchFunc = func(c context.Context, K dataloader.Keys) (out []*dataloader.Result) {
	out = []*dataloader.Result{
		{Data: "The quick brown fox jumps over the lazy dog."},
	}

	return
}

func TestDataLoaderNewBatchLoader(t *testing.T) {
	ldr := dataloader.NewBatchedLoader(exampleBatchFn)

	result, err := ldr.Load(context.TODO(), dataloader.StringKey(""))()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%s\n", result.(string))
}
