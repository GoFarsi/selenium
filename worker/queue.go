package worker

import "sync"

type queue struct {
	lk        *sync.Mutex
	proxyList []string
}

func newQueue(proxyList []string) *queue {
	return &queue{
		lk:        &sync.Mutex{},
		proxyList: proxyList,
	}
}

func (q *queue) pop(proxy string) {
	q.lk.Lock()
	defer q.lk.Unlock()
	for i, v := range q.proxyList {
		if v == proxy {
			q.proxyList[i] = q.proxyList[len(q.proxyList)-1]
			q.proxyList = q.proxyList[:len(q.proxyList)-1]
			break
		}
	}
}
