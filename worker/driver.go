package worker

import (
	"context"
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type Result struct {
	Target   string
	Proxy    string
	Title    string
	WorkerId int
	Err      error
}

func startView(target string, proxyList []string, numOfWorker int, seleniumServerPath, driverPath string, debug bool, result chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(proxyList)/numOfWorker; i++ {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		queue := newQueue(proxyList)
		wgWorker := &sync.WaitGroup{}
		wgWorker.Add(numOfWorker)
		for j := 0; j < numOfWorker; j++ {
			go func() {
				result <- startDriver(ctx, i, target, queue.getProxy(), seleniumServerPath, driverPath, debug, wgWorker)
			}()
		}
		wgWorker.Wait()
	}
}

func startDriver(ctx context.Context, workerId int, target string, proxy string, seleniumServerPath, driverPath string, debug bool, wg *sync.WaitGroup) Result {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return Result{target, proxy, "", workerId, ctx.Err()}
	default:
	}
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),
		selenium.ChromeDriver(driverPath),
	}
	if debug {
		selenium.SetDebug(true)
		opts = append(opts, selenium.Output(os.Stdout))
	} else {
		opts = append(opts, selenium.Output(nil))
	}
	port, err := getFreePort()
	if err != nil {
		return Result{target, proxy, "", workerId, err}
	}
	service, err := selenium.NewSeleniumService(seleniumServerPath, port, opts...)
	if err != nil {
		return Result{target, proxy, "", workerId, err}
	}
	defer service.Stop()
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddProxy(selenium.Proxy{
		Type: selenium.Manual,
		HTTP: proxy,
	})
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return Result{target, proxy, "", workerId, err}
	}
	defer wd.Quit()
	log.Printf("worker %d : task view site %s on proxy %s has been started", workerId, target, proxy)
	if err := wd.Get(target); err != nil {
		return Result{target, proxy, "", workerId, err}
	}
	title, err := wd.Title()
	return Result{target, proxy, title, workerId, err}
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
