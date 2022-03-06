package helpers

type IssueResult struct {
	TotalCount       int  `json:"total_count"`
	IncompleteResult bool `json:"incomplete_result"`
	//Items            []*Items
}

//type IssuesSearchResult struct {
//	TotalCount int `json:"total_count"`
//	Items      []*Items
//}
//
//type Items struct {
//	Number    int
//	HTMLURL   string `json:"html_url"`
//	Title     string
//	State     string
//	User      *User
//	CreatedAt time.Time `json:"created_at"`
//	Body      string    // in Markdown format
//}
//
//type User struct {
//	Login   string
//	HTMLURL string `json:"html_url"`
//}
