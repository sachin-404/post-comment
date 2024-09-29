package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"-" gorm:"not null"`
}

func NewUser(name, email, password string) *User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
