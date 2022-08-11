package worker

import (
	"context"
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
	res := make(chan Result)
	for _, tt := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		go startDriver(ctx, tt.Target, tt.Proxy, seleniumPath, chromePath, res, false)
	}
	for result := range res {
		if result.Err != nil {
			t.Fatalf("Target %v get error %v", result.Target, result.Err)
		}
		t.Logf("Target %v Title %v viewed with Proxy %v", result.Target, result.Title, result.Proxy)
	}
	close(res)
}
