package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func cronFn() {
	fmt.Println("Cron is running")
}

func main() {
	s := gocron.NewScheduler()
	// s.Every(2).Hours().Do(task)
	s.Every(1).Day().At("16:50").Do(cronFn)
	<-s.Start()
}
