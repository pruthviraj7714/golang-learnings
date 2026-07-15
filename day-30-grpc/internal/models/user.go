package models

type User struct {
	ID       int64  `json:"id" gorm:"autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
