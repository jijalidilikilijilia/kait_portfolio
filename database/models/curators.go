package models

type Curators struct {
	Id        uint   `gorm:"primary key;autoIncrement" json:"id"`
	Full_name string `json:"full_name"`
	Phone     string `json:"phone"`
}
