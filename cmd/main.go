package main

import "github.com/romv7/blogs/internal/utils"

func main() {
	if err := utils.NewSQLSchemaTool().ReloadAll(); err != nil {
		println(err.Error())
	}
}
