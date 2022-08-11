package worker

import "testing"

func TestStartDriver(t *testing.T) {
	target, proxy := "https://google.com", "41.65.236.44:1976"
	if err := startDriver(target, proxy); err != nil {
		t.Error(err)
	}
}
