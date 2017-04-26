package main

import (
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/pkg/errors"
)

// getSources returns a list of possible feeds
// You should use sourcesURL url.
func getSources(url string) ([]*feed, error) {
	body, err := doHTTP(url)
	if err != nil {
		return nil, errors.New("getFeeds - " + err.Error())
	}

	list := make([]*feed, 0)

	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, errors.New("Wrong Sources Json - " + err.Error())
	}
	return list, nil
}

// getNews returns a list of news.
// You should use mainFeedURL url.
func getNews(url string) ([]*news, error) {
	body, err := doHTTP(url)
	if err != nil {
		return nil, errors.New("getNews - " + err.Error())
	}

	list := make([]*news, 0)

	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, errors.New("Wrong News Json - " + err.Error())
	}
	return list, nil
}

// doHTTP makes a simple GET http call and returns the content
func doHTTP(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Get - " + err.Error())
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ReadAll - " + err.Error())
	}
	resp.Body.Close()

	return respBody, nil

}
