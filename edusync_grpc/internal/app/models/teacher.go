package models

type Teacher struct {
	Model
	Name         string      `json:"name"`
	Email        string      `json:"email" gorm:"uniqueIndex"`
	PhoneNumber  string      `json:"phone_number"`
	Courses      []*Course   `json:"courses" gorm:"many2many:course_teachers;"`
	DepartmentID *string     `json:"department_id"`
	Department   *Department `json:"department" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AddressID    *string     `json:"address_id"`
	Address      *Address    `json:"address" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
