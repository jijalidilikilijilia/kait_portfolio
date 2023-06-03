package models

type StudentWorks struct {
	Id          uint   `gorm:"primary key;autoIncrement" json:"id"`
	Student_id  uint   `gorm:"student_id"`
	File_name   string `gorm:"file_name"`
	Upload_date string `gorm:"type:date;default:current_date"`
	File        []byte `gorm:"file"`
}
