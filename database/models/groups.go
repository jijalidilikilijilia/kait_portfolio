package models

type Groups struct {
	Id         uint   `gorm:"primary key;autoIncrement" json:"id"`
	Group_name string `json:"group_name"`
	Curator_id int    `json:"curator_id"`
}
