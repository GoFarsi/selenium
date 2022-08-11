package worker

import (
	"context"
	"fmt"
	"github.com/tebeka/selenium"
	"os"
)

type result struct {
	target string
	title  string
	err    error
}

func startDriver(ctx context.Context, target string, proxy string, seleniumServerPath, driverPath string, port int, res chan result) {
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),
		selenium.ChromeDriver(driverPath),
		selenium.Output(os.Stderr),
	}
	service, err := selenium.NewSeleniumService(seleniumServerPath, port, opts...)
	if err != nil {
		res <- result{target, "", err}
	}
	defer service.Stop()
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddProxy(selenium.Proxy{
		Type: selenium.Manual,
		HTTP: proxy,
	})
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		res <- result{target, "", err}
	}
	defer wd.Quit()
	if err := wd.Get(target); err != nil {
		res <- result{target, "", err}
	}
	title, err := wd.Title()
	res <- result{target, title, err}

	for {
		select {
		case <-ctx.Done():
			res <- result{target, "", ctx.Err()}
		case <-res:
		}
	}

}
