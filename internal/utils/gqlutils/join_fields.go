package gqlutils

import (
	"errors"
	"log"

	gql "github.com/graphql-go/graphql"
)

var (
	ErrFieldsHasConflictName = errors.New("attempting to join fields that has conflicting keys.")
)

// Joins a set of gq.Fields into a flat one. Take note, if there are fields
// that has the same key as the one that occured in the first insertion,
// it will cause a runtime error.
func JoinFields(fs ...gql.Fields) (f gql.Fields) {
	exists := make(map[string]bool)
	f = make(gql.Fields)

	for _, fields := range fs {

		for k, field := range fields {
			if exists[k] {
				log.Panic(ErrFieldsHasConflictName)
			}

			f[k] = field
			exists[k] = true
		}

	}

	return
}

// Adds the set of gql.Fields into an object's fields map.
func AddFields(o *gql.Object, fs ...gql.Fields) {

	for _, fields := range fs {
		for k, field := range fields {
			o.AddFieldConfig(k, field)
		}
	}

}
