package main

import (
	"fmt"
	
	"github.com/robfig/cron/v3"
)

func Hello(t string) {
	fmt.Println("hello " + t)
}

func EnableScheduler() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {
		go Hello("new world order")
	})
	c.Start()
}