package model

type User struct {
	Name string
}

type Link struct {
	Title string `json:"title"`
	Address string `json:"address"`
	User User `json:"user"`
}