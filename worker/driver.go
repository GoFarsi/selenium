package worker

import (
	"fmt"
	"github.com/tebeka/selenium"
	"os"
	"time"
)

const (
	seleniumServerPath = "/home/user/Project/go/Personal/selenium/download/selenium-server.jar"
	driverPath         = "/home/user/Project/go/Personal/selenium/download/chromedriver"
	port               = 8080
)

func startDriver(target string, proxy string) error {
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),
		selenium.ChromeDriver(driverPath),
		selenium.Output(os.Stderr),
	}
	service, err := selenium.NewSeleniumService(seleniumServerPath, port, opts...)
	if err != nil {
		return err
	}
	defer service.Stop()
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddProxy(selenium.Proxy{
		Type: selenium.Manual,
		HTTP: proxy,
	})
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return err
	}
	defer wd.Quit()
	if err := wd.Get(target); err != nil {
		return err
	}
	if err := wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		return false, nil
	}, 3*time.Second); err != nil {
		return err
	}
	return nil
}
