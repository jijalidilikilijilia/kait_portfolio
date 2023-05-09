package models

type Cumpus struct {
	Id          int    `gorm:"primary key;autoIncrement" json:"id"`
	Cumpus_name string `json:"cumpus_name"`
}
