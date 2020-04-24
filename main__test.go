package main

import (
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

func doRequest(uri string) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func TestGet(t *testing.T) {
	N := 1000
	wg := sync.WaitGroup{}
	wg.Add(N)

	start := time.Now()
	for i := 0; i < N; i++ {
		go func() {
			if err := doRequest("https://kreasindo-web.herokuapp.com/"); err != nil {
				t.Error(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	t.Logf("Total duration for %d concurrent request(s) is %v", N, time.Since(start))
}
