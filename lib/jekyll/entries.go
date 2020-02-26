package jekyll

type Pages struct {
	Entries []PageEntry `json:"entries"`
}

type PageEntry struct {
	Title string   `json:"title"`
	Url   string   `json:"url"`
	Tags  []string `json:"tags"`
	Body  string   `json:"body"`
}
