package models

type Student struct {
	Id          uint   `gorm:"primary key;autoIncrement" json:"id"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Full_name   string `json:"full_name"`
	Age         int    `json:"age"`
	Group_id    int    `json:"group_id"`
	Speciality  int    `json:"speciality_id"`
	Description string `json:"desc"`
}
