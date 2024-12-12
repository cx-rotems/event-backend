package models

type Name struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Arrived   bool   `json:"arrived"`
}
