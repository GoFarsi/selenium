package cmd

import "github.com/alexflint/go-arg"

type Args struct {
	Address            []string `arg:"-a" help:"list of address for view for example -a http://example1.com http://example2.com"`
	ProxyPath          string   `arg:"-p" help:"address of path proxy list for example -p /home/user/proxy.txt"`
	NumOfWorkers       uint     `arg:"-w" default:"2" help:"set number of worker for concurrency view in same time"`
	SeleniumServerPath string   `arg:"-s" help:"path of selenium server file for example /home/user/selenium-server.jar"`
	ChromeDriverPath   string   `arg:"-c" help:" path of chrome driver for example /home/user/chromedriver"`
}

func InitCommands() *Args {
	args := &Args{}
	arg.MustParse(args)
	return args
}
