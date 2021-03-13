package models

type Domains struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Domain string `json:"domain"`
}
