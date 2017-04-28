package main

import (
	"strings"
	"time"
)

type (
	// feed is a struct for sourcesURL
	feed struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Image       string `json:"image"`
		Category    string `json:"category"`
		Link        string `json:"link"`
		Description string `json:"description"`
	}

	// news is a struct for mainFeedURL
	news struct {
		Feed      string  `json:"feed"`
		Icon      string  `json:"favicon"`
		Image     string  `json:"image"`
		Type      string  `json:"type"`
		Proof     int     `json:"-"`
		Timestamp *myTime `json:"timestamp"`
		URL       string  `json:"url"`
		Title     string  `json:"title"`
		Category  string  `json:"category"`
	}

	myTime struct {
		time *time.Time
	}
)

func (t *myTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	dataSplit := strings.Split(string(data), ".")

	parsedTime, err := time.Parse(`"`+"2006-01-02 15:04:05"+`"`, dataSplit[0]+`"`)
	if err != nil {
		return err
	}

	t.time = &parsedTime

	return nil
}

func (t *myTime) String() string {
	if t.time != nil {
		return (t.time).Format("2006-01-02 - 15:04")
	}
	return ""
}
