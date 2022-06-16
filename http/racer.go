package http

import (
	"fmt"
	"net/http"
	"time"
)

type TimeoutErr struct {
	error
	urls []string
}

func (e TimeoutErr) Error() string {
	return fmt.Sprintf("timed out after waiting for %s", e.urls)
}

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (fastestUrl string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", TimeoutErr{urls: []string{a, b}}
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
