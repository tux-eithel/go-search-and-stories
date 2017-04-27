package main

import "net/url"
import "strings"

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
