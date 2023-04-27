package main

// https://betterprogramming.pub/parsing-and-creating-yaml-in-go-crash-course-2ec10b7db850

import (
	"fmt"
	"os"
	"restic-bot/util"
)

func main() {
	f, err := os.ReadFile("test.yml")
	util.CheckError(err)

	var conf = util.ParseYml(f)

	fmt.Printf("%+v\n", conf)
}
