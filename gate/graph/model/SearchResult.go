package model

type SearchResult struct {
	Page      int     `json:"page"`
	Total     int     `json:"total"`
	Result    []*Item `json:"result"`
	ResultIds []int   `json:"resultIds"`
}
