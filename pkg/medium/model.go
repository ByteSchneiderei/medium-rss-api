package medium

// Response represents blog api response body
type Response struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Posts       []Post `json:"posts"`
}

// Post represents single post in blog api response body
type Post struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Link       string   `json:"link"`
	Published  string   `json:"published"`
	Author     string   `json:"author"`
	Categories []string `json:"categories"`
	Content    *Node    `json:"content"`
}

// Node represents DOM element in HTML
type Node struct {
	Tag      string              `json:"tag"`
	Text     string              `json:"text"`
	Attrs    []map[string]string `json:"attrs"`
	Children []*Node             `json:"children"`
}
