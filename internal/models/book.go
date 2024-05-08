package models

type Book struct {
	Id             int    `json:"-"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ProductionYear int    `json:"production_year"`
}
