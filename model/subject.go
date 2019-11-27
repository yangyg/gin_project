package model

import "github.com/jinzhu/gorm"

// Subject model
type Subject struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

// SubjectCreate create
func (subject *Subject) SubjectCreate() (uint, error) {
	result := Db.Create(subject)

	return subject.ID, result.Error
}

// UpdateSubject .
func (subject *Subject) UpdateSubject() {
	Db.Update(subject)
}

// DeleteSubject .
func (subject *Subject) DeleteSubject() {
	Db.Delete(subject)
}

// SubjectList ...
func SubjectList() ([]Subject, error) {
	subjects := make([]Subject, 0)
	result := Db.Find(&subjects)

	return subjects, result.Error
}
