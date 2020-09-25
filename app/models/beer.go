package models

type Beer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"typeBeer"`
	Brewery string `json:"brewery"`
}
