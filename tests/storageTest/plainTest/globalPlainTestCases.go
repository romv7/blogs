package plainTest

import (
	"os"
)

type plainTestCase struct {
	rootPath string
	key      string
	content  []byte
}

type plainTestCases = []*plainTestCase

var (
	rootPath = os.Getenv("STORAGE_DIR")

	globalPlainTestCases = plainTestCases{
		{
			rootPath: rootPath,
			key:      "hello.txt",
			content:  []byte("The quick brown fox jumps over the lazy dog."),
		},
		{
			rootPath: rootPath,
			key:      "blogs/tests/message.txt",
			content:  []byte("Do you think I'm a good person?"),
		},
		{
			rootPath: rootPath,
			key:      "gotests/main.go",
			content: []byte(`
			package main

			import "fmt"

			func main() {
				fmt.Println("hello, world")
			}
			`),
		},
	}
)
