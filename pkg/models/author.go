package models

type AuthorDetail struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	AuthorName  string
	Address     string
	PhoneNumber string
}
