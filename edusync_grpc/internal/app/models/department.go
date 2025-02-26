package models

type Department struct {
	Model
	Name        string     `json:"name" gorm:"index"`
	Email       string     `json:"email" gorm:"uniqueIndex"`
	PhoneNumber string     `json:"phone_number"`
	AddressID   *string    `json:"address_id"`
	Address     *Address   `json:"address" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Teachers    []*Teacher `json:"teachers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Students    []*Student `json:"students" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Courses     []*Course  `json:"courses" gorm:"many2many:department_courses;"`
}
