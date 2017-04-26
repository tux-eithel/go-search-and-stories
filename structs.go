package main

type (
	feed struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Image    string `json:"image"`
		Category string `json:"category"`
		Link     string `json:"link"`
	}

	news struct {
		Feed     string `json:"feed"`
		Icon     string `json:"favicon"`
		Image    string `json:"image"`
		Type     string `json:"type"`
		Proof    int    `json:"proof"`
		URL      string `json:"url"`
		Title    string `json:"title"`
		Category string `json:"category"`
	}
)
