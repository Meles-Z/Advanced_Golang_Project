package models

type Address struct {
	Model
	Street   string `json:"street"`
	Location string `json:"location"`
	Region   string `json:"region"`
	Country  string `json:"country"`
}
