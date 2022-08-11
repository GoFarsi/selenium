package worker

import "testing"

func TestQueuePop(t *testing.T) {
	proxyList := []string{
		"192.111.135.18",
		"202.183.24.38",
		"62.201.252.158",
		"93.184.151.48",
		"181.143.69.227",
		"209.97.176.112",
		"79.132.221.205",
		"192.111.129.145",
		"193.84.184.25",
		"202.179.184.46",
	}
	queue := newQueue(proxyList)
	for _, proxy := range proxyList {
		queue.pop(proxy)
		t.Log(queue.proxyList)
	}
}
