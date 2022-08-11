package main

import (
	"github.com/Ja7ad/selenium/cmd"
	"github.com/Ja7ad/selenium/worker"
	"log"
	"os"
)

func main() {
	c := cmd.InitCommands()
	w := worker.NewWorker(c.Address, c.ProxyPath, c.SeleniumServerPath, c.ChromeDriverPath, c.NumOfWorkers, c.Debug)
	result, err := w.Start()
	if err != nil {
		log.Fatal(err)
	}
	for res := range result {
		if res.Err != nil {
			log.Printf("worker %d : task view site %s go error %v on proxy %s", res.WorkerId, res.Target, res.Err, res.Proxy)
			continue
		}
		log.Printf("worker %d : task view site %s with title %s on proxy %s has been done", res.WorkerId, res.Target, res.Title, res.Proxy)
	}
	os.Exit(0)
}
