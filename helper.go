package main

import "net/url"
import "strings"
import "os"
import "log"
import "io/ioutil"
import "encoding/json"

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

func loadSettings(filname string) *preference {
	file, err := os.Open(filname)
	if err != nil {
		log.Println("Unable to open setting file", filname)
		return &preference{}
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Unable to read setting file", err)
		return &preference{}
	}

	p := &preference{}
	json.Unmarshal(data, p)
	if err != nil {
		log.Println("Unable to parse the json", err)
		return &preference{}
	}
	return p

}

func saveSettings(filename string, p *preference) bool {

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
