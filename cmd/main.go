package main

import "github.com/rommms07/blogs/internal/utils"

func main() {
	if err := utils.NewSQLSchemaTool().ReloadAll(); err != nil {
		println(err.Error())
	}
}
