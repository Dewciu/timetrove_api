package models

type Address struct {
	BaseModel
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}
