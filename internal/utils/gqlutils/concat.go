// This utility is used for recursively concatenating the schema definition files (SDL)
// declared in the endpoints/gql/schemas directory.
package gqlutils

import (
	"bytes"
	"io/fs"
	"log"
	"os"
	"strings"
)

var (
	rootGqlDir = os.Getenv("ROOTDIR") + "/endpoints/gql"
)

func JoinAllSchemaFiles() (out string) {
	rootfs := os.DirFS(rootGqlDir)

	var buf bytes.Buffer

	var err error
	var walkfn fs.WalkDirFunc = func(path string, d fs.DirEntry, err error) error {
		if !(strings.HasSuffix(path, ".graphql") || strings.HasSuffix(path, ".gql")) {
			return nil
		}

		b, err := os.ReadFile(rootGqlDir + "/" + path)
		if err != nil {
			return err
		}

		if _, err := buf.Write(append(b, '\n')); err != nil {
			return err
		}

		return nil
	}

	err = fs.WalkDir(rootfs, ".", walkfn)
	if err != nil {
		log.Panic(err)
	}

	return buf.String()
}
