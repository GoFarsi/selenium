package worker

import (
	"context"
	"fmt"
	"github.com/tebeka/selenium"
	"net"
	"sync"
)

type result struct {
	target string
	proxy  string
	title  string
	err    error
}

func startDriver(ctx context.Context, target string, proxy string, seleniumServerPath, driverPath string, res chan result, wg *sync.WaitGroup) {
	defer wg.Done()
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),
		selenium.ChromeDriver(driverPath),
		selenium.Output(nil),
	}
	port, err := getFreePort()
	if err != nil {
		res <- result{target, proxy, "", err}
	}
	service, err := selenium.NewSeleniumService(seleniumServerPath, port, opts...)
	if err != nil {
		res <- result{target, proxy, "", err}
	}
	defer service.Stop()
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddProxy(selenium.Proxy{
		Type: selenium.Manual,
		HTTP: proxy,
	})
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		res <- result{target, proxy, "", err}
	}
	defer wd.Quit()
	if err := wd.Get(target); err != nil {
		res <- result{target, proxy, "", err}
	}
	title, err := wd.Title()
	res <- result{target, proxy, title, err}

	for {
		select {
		case <-ctx.Done():
			res <- result{target, proxy, "", ctx.Err()}
		case <-res:
		}
	}
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
