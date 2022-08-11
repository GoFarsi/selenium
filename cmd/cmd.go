package cmd

import "github.com/alexflint/go-arg"

type Args struct {
	Address            []string `arg:"-a, required" help:"list of address for view for example -a http://example1.com http://example2.com"`
	ProxyPath          string   `arg:"-p, required" help:"address of path proxy list for example -p /home/user/proxy.txt"`
	NumOfWorkers       int      `arg:"-w" default:"2" help:"set number of worker for concurrency view in same time"`
	SeleniumServerPath string   `arg:"-s, required" help:"path of selenium server file for example /home/user/selenium-server.jar"`
	ChromeDriverPath   string   `arg:"-c, required" help:" path of chrome driver for example /home/user/chromedriver"`
	Debug              bool     `arg:"env:DEBUG" default:"false" help:"show debug logs of selenium"`
}

func InitCommands() *Args {
	args := &Args{}
	arg.MustParse(args)
	return args
}
