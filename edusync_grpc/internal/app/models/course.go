package models

type Course struct {
	Model
	Name     string     `json:"name" gorm:"index"`
	Teachers []*Teacher `json:"teachers" gorm:"many2many:course_teachers;"`
	Students []*Student `json:"students" gorm:"many2many:course_students;"`
}
