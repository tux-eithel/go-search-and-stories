package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/user"
	"sort"
	"strings"
)

// filterSources add listSources ids to query string.
// The ids will be added to "s" param.
// If listSources is nil, mainFeedURL url will be returned
func filterSources(inputURL string, listSources []string) *url.URL {

	currentURL, err := url.Parse(inputURL)
	if err != nil {
		return nil
	}

	if listSources == nil || len(listSources) == 0 {
		return currentURL
	}

	params := currentURL.Query()
	params.Add("s", strings.Join(listSources, ","))
	currentURL.RawQuery = params.Encode()
	return currentURL
}

func feedByCategory(feeds []*feed) map[string][]*feed {

	orderedFeed := make(map[string][]*feed, 0)

	for _, currentFeed := range feeds {
		if _, ok := orderedFeed[currentFeed.Category]; !ok {
			orderedFeed[currentFeed.Category] = make([]*feed, 0)
		}
		orderedFeed[currentFeed.Category] = append(orderedFeed[currentFeed.Category], currentFeed)
	}

	return orderedFeed

}

func orderFeedByName(feeds map[string][]*feed) map[string][]*feed {

	sortedFeeds := make(map[string][]*feed, len(feeds))

	for index, value := range feeds {
		sort.Slice(value, func(i, j int) bool {
			return value[i].Title < value[j].Title
		})
		sortedFeeds[index] = value
	}

	return sortedFeeds

}

func loadSettings(filename string) *preference {

	p := &preference{}

	user, err := user.Current()
	if err != nil {
		log.Println("Unable to determinate the home directory", err)
		return p
	}

	filename = strings.Join([]string{user.HomeDir, filename}, string(os.PathSeparator))

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Unable to open setting file", filename)
		return p
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Unable to read setting file", err)
		return p
	}

	json.Unmarshal(data, p)
	if err != nil {
		log.Println("Unable to parse the json", err)
		return p
	}
	return p

}

func saveSettings(filename string, p *preference) bool {

	user, err := user.Current()
	if err != nil {
		log.Println("Unable to determinate the home directory", err)
		return false
	}

	filename = strings.Join([]string{user.HomeDir, filename}, string(os.PathSeparator))

	data, err := json.Marshal(p)
	if err != nil {
		log.Println("Unable to marshal settings", err)
		return false
	}

	err = ioutil.WriteFile(filename, data, 0600)
	if err != nil {
		log.Println("Unable to save the settings", err)
		return false
	}
	return true

}
