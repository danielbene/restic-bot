package main

import (
	"fmt"
	"os"
	"restic-bot/util"
)

func main() {
	f, err := os.ReadFile("config.yml")
	util.CheckError(err)

	var conf = *util.ParseYml(f)

	fmt.Printf("%+v\n", conf)
	fmt.Println(conf.Look)
}
