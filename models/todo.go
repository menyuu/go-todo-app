package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string `form:"title" binding:"required,max=50"`
	Done   bool
	UserID uint
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
func CreateTodo(title string) error {
	todo := Todo{Title: title, Done: false}
	return DB.Create(&todo).Error
}

// 変更
func UpdateTodo(id int, title string, done bool) error {
	var todo Todo
	if err := DB.First(&todo, id).Error; err != nil {
		return err
	}

	todo.Title = title
	todo.Done = done
	return DB.Save(&todo).Error
}

// 削除
func DeleteTodo(id int) error {
	return DB.Delete(&Todo{}, id).Error
}
