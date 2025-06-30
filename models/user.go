package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string `json:"-"`
}

// 全件取得
func GetAllUsers() ([]User, error) {
	var users []User
	err := DB.Find(&users).Error
	return users, err
}

// 1件取得
func GetUserByID(id int) (User, error) {
	var user User
	err := DB.First(&user, id).Error
	return user, err
}

// 作成
func CreateUser(user *User) error {
	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashPassword

	return DB.Create(&user).Error
}

// 変更
func UpdateUser(user User) error {
	if err := DB.First(&user, user.ID).Error; err != nil {
		return err
	}

	if user.Password != "" {
		hashPassword, err := HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashPassword
	}

	return DB.Save(&user).Error
}

// 削除
func DeleteUser(id string) error {
	return DB.Delete(&User{}, id).Error
}

// パスワードのハッシュ化
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func AuthenticateUser(email, password string) (*User, error) {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
