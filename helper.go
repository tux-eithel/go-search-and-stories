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

	if listSources == nil {
		return currentURL
	}

	params := currentURL.Query()
	params.Add("s", strings.Join(listSources, ","))
	currentURL.RawQuery = params.Encode()
	return currentURL
}
