package models

type Specialities struct {
	Id              uint   `gorm:"primary key;autoIncrement" json:"id"`
	Speciality_name string `json:"speciality_name"`
}
