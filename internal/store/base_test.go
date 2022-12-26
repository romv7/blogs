package store_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/rommms07/blogs/internal/store"
)

type mockDataSource struct {
}

func (s *mockDataSource) Connect() (*mockDataSource, error) {
	return s, errors.New("error: unimplemented data source...")
}

func Test_shallowCheckUnimplementedStubs(t *testing.T) {
	unimp := &store.UnimplementedStore{}
	expectedCases := map[string]func() error{
		"UnimplementedNew":    unimp.New,
		"UnimplementedSave":   unimp.Save,
		"UnimplementedDelete": unimp.Delete,
		"UnimplementedRead":   unimp.Read,
	}

	for caseName, unimpMethod := range expectedCases {
		if err := unimpMethod(); err == nil {
			t.Errorf("[fail] %s did not returned an error..", caseName)
		}
	}
}

func Test_creatingNewDataSourceFromMockSource(t *testing.T) {
	src, err := store.NewDataSource(func(db *store.DataSource[mockDataSource]) {
		db.Source = &mockDataSource{}
	})

	if err != nil {
		t.Logf(err.Error())
	}

	sameSrc, err := src.Connect()

	if err == nil {
		t.Errorf("[fail] src.Connect is expected to return an unimplemented error...")
	}

	if fmt.Sprintf("%x", src.Source) != fmt.Sprintf("%x", sameSrc) {
		t.Errorf("[fail] src.Connect did not returned the same expected data source...")
	}
}

func Test_expectedNewDataSourceToReturnAnErrorWhenNotInitializedProperly(t *testing.T) {
	_, err := store.NewDataSource(func(db *store.DataSource[mockDataSource]) {})

	if err == nil {
		t.Errorf("[fail] store.NewDataSource inapproriately instantiate the data store...")
	}
}
