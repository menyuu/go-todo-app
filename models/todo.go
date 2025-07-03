package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string
	Done   bool
	UserID uint
	User   User `gorm:"constraint:OnDelete:CASCADE;"`
}

// 全件取得
func GetAllTodos() ([]Todo, error) {
	var todos []Todo
	if err := DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// 1件取得
func GetTodoByID(id int) (*Todo, error) {
	var todo Todo
	if err := DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// 作成
func CreateTodo(todo *Todo) error {
	return DB.Create(&todo).Error
}

// 変更
func UpdateTodo(todo *Todo) error {
	if err := DB.First(&todo).Error; err != nil {
		return err
	}

	return DB.Save(&todo).Error
}

// 削除
func DeleteTodo(id int) error {
	return DB.Delete(&Todo{}, id).Error
}
