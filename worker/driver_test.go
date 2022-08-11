package worker

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestStartDriver(t *testing.T) {
	seleniumPath := "/home/user/Project/go/Personal/selenium/download/selenium-server.jar"
	chromePath := "/home/user/Project/go/Personal/selenium/download/chromedriver"
	tests := []struct {
		Target string
		Proxy  string
	}{
		{"https://google.com", "41.65.236.44:1976"},
		{"https://bing.com", "41.65.236.44:1976"},
		{"https://soft98.ir", "41.65.236.44:1976"},
		{"https://bscscan.com/token/0xacfc95585d80ab62f67a14c566c1b7a49fe91167", "41.65.236.44:1976"},
		{"https://www.coingecko.com/en/coins/feg-token-bsc", "41.65.236.44:1976"},
	}
	res := make(chan result)
	wg := &sync.WaitGroup{}
	wg.Add(len(tests))
	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		go startDriver(ctx, tt.Target, tt.Proxy, seleniumPath, chromePath, res, wg)
	}
	for result := range res {
		if result.err != nil {
			t.Fatalf("target %v get error %v", result.target, result.err)
		}
		t.Logf("target %v title %v viewed with proxy %v", result.target, result.title, result.proxy)
	}
	close(res)
	wg.Wait()
}
