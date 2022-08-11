package worker

import "testing"

func TestWorker(t *testing.T) {
	test := []*Worker{
		{
			Address:     []string{"https://example1.com", "https://example2.com"},
			ProxyPath:   "test/Proxy.txt",
			NumOfWorker: 3,
		},
		{
			Address:     []string{},
			ProxyPath:   "test/Proxy.txt",
			NumOfWorker: 3,
		},
		{
			Address:     []string{"https://example1.com", "https://example2.com"},
			ProxyPath:   "",
			NumOfWorker: 3,
		},
		{
			Address:     []string{"https://example1.com", "https://example2.com"},
			ProxyPath:   "test/Proxy.txt",
			NumOfWorker: 5,
		},
		{
			Address:     []string{"https://example1.com", "https://example2.com"},
			ProxyPath:   "test/Proxy.txt",
			NumOfWorker: 0,
		},
	}

	for _, tt := range test {
		t.Run("", func(t *testing.T) {
			if err := tt.checkAddress(); err != nil {
				t.Error(err)
			}
			if err := tt.setProxyListFromFile(); err != nil {
				t.Error(err)
			}
			if err := tt.checkNumOfWorker(); err != nil {
				t.Error(err)
			}
			t.Log(tt)
		})
	}
}
