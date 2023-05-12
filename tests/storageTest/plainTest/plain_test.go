package plainTest

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/romv7/blogs/internal/storage/driver/plain"
)

func TestNewPlain(t *testing.T) {
	if plain.Default == nil {
		t.Error("plain.Default must not be nil")
	}
}

func TestShouldGetFileUsingTheKey(t *testing.T) {
	for _, tcase := range globalPlainTestCases {
		fileName := strings.Join([]string{tcase.rootPath, tcase.key}, "/")

		// Create the directory described in the `fileName`
		tokens := strings.Split(strings.Replace(fileName, tcase.rootPath, "", 1), "/")[1:]

		for i := range tokens[:len(tokens)-1] {
			path := []string{tcase.rootPath}

			if i > 0 {
				path = append(path, tokens[i-1:len(tokens)-1]...)
			} else {
				path = append(path, tokens[i])
			}

			p := strings.Join(path, "/")

			os.Mkdir(p, os.FileMode(0700))
			defer os.Remove(p)
		}

		if err := os.WriteFile(fileName, tcase.content, os.FileMode(0644)); err != nil {
			log.Panic(err)
		}

		defer os.Remove(fileName)

		if p, err := plain.Default.Get(tcase.key); err != nil {
			t.Error(err)
		} else if bytes.Compare(tcase.content, p) != 0 {
			t.Errorf("content did not matched")
		}

	}
}

func Test_ShouldPutFileToThePathSpecifiedByKey(t *testing.T) {
	for _, tcase := range globalPlainTestCases {

		if err := plain.Default.Put(tcase.key, tcase.content); err != nil {
			t.Error(err)
		}

		defer os.Remove(strings.Join([]string{tcase.rootPath, tcase.key}, "/"))
	}
}
