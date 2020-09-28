package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

type Provider struct {
	index int
	mux sync.Mutex
}

var wg = &sync.WaitGroup{}

func (provider * Provider) getIndex() int {
	provider.mux.Lock()
	defer provider.mux.Unlock()

	if provider.index == -1 {
		return -1
	}

	provider.index += 1
	return provider.index - 1
}

func (provider * Provider) setIndex(value int) {
	provider.mux.Lock()
	defer provider.mux.Unlock()

	provider.index = value
}

func main() {
	var provider = Provider{
		index: 0,
		mux:   sync.Mutex{},
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go WorkThread(&provider, "http://name-service.appspot.com/api/v1/names/")
	}

	wg.Wait()
}

func WorkThread(provider * Provider, url string) {
	defer wg.Done()
	for fileIndex := provider.getIndex(); fileIndex != -1; fileIndex = provider.getIndex() {
		err := DownloadFile("./data/" + strconv.Itoa(fileIndex) + ".json", url + strconv.Itoa(fileIndex) + ".json")
		if err != nil {
			provider.setIndex(-1)
			return
		}
	}
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	isJson := json.Valid(body)
	if isJson != true {
		return http.ErrNotSupported
	}

	err = ioutil.WriteFile(filepath, body, 0600)
	if err != nil {
		return err
	}

	return nil
}
