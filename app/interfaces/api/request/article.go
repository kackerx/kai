package request

type EditArticleReq struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
