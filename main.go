package main

import (
	"github.com/Ja7ad/selenium/cmd"
	"github.com/Ja7ad/selenium/worker"
	"log"
)

func main() {
	c := cmd.InitCommands()
	worker := worker.NewWorker(c.Address, c.ProxyPath, c.NumOfWorkers)
	if err := worker.Start(); err != nil {
		log.Fatal(err)
	}
}
