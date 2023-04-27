package main

import (
	"fmt"
	"os"
	"restic-bot/util"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	f, err := os.ReadFile("config.yml")
	util.CheckError(err)

	var conf = *util.ParseYml(f)

	fmt.Printf("%+v\n", conf)
	fmt.Println(conf.Look)

	cron := gocron.NewScheduler(time.Local)

	cron.Every(5).Seconds().Do(func() {
		fmt.Println("yo")
	})

	cron.Every("8s").Do(func() {
		fmt.Println("hai")
	})

	cron.StartBlocking()
}
