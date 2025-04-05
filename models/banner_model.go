package models

type Banner struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Src  string  `json:"src"`
	Link *string `json:"link"`
}
