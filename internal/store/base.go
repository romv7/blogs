package store

import (
	"errors"

	"github.com/romv7/blogs/internal/store/source"
	"github.com/romv7/blogs/internal/stub"
)

type UnimplementedStore struct{}

func (s *UnimplementedStore) New() error {
	return stub.UnimplementedMethodCalled()
}

func (s *UnimplementedStore) Save() error {
	return stub.UnimplementedMethodCalled()
}

func (s *UnimplementedStore) DeleteByUuid() error {
	return stub.UnimplementedMethodCalled()
}

func (s *UnimplementedStore) DeleteById() error {
	return stub.UnimplementedMethodCalled()
}

func (s *UnimplementedStore) Read() error {
	return stub.UnimplementedMethodCalled()
}

type DataSource[T any] struct {
	IsValid bool
	Source  source.SourceConnector[T]
}

func NewDataSource[T any](initCb func(*DataSource[T])) (*DataSource[T], error) {
	src := &DataSource[T]{IsValid: true}
	initCb(src)

	if src.Source == nil {
		return nil, errors.New("error: callback did not properly initialized the data source...")
	}

	return src, nil
}

func (s *DataSource[T]) Connect() (*T, error) {
	src, err := s.Source.Connect()
	if err != nil {
		s.IsValid = false
		return src, err
	}
	return src, err
}
