package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique" json:"email" validate:"required,email"`
	Password  string `json:"-" validate:"required,min=4,max=32"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPasswordHash(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return (err == nil), err
}
