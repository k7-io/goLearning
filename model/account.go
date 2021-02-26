package model

type Account struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}
