package model

type Entry struct {
	Sys struct {
		ID        string `json:"id"`
		CreatedAt string `json:"createdAt"`
	} `json:"sys"`
	Fields struct {
		Name string `json:"name"`
	} `json:"fields"`
}
