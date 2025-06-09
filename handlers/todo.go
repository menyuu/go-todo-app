package handlers

import (
	"fmt"
	"go-todo-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET 一覧
func Index(c *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		c.String(http.StatusInternalServerError, "ToDoの取得に失敗しました")
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"todos": todos,
	})
}

func Show(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "IDが不正です")
		return
	}

	todo, err := models.GetTodoByID(id)
	if err != nil {
		c.String(http.StatusNotFound, "ToDoが見つかりません")
		return
	}

	c.HTML(http.StatusOK, "show", gin.H{"todo": todo})
}

// GET 作成
func New(c *gin.Context) {
	c.HTML(http.StatusOK, "new", nil)
}

// GET 変更
func Edit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "IDが不正です")
		return
	}

	todo, err := models.GetTodoByID(id)
	if err != nil {
		c.String(http.StatusNotFound, "ToDoが見つかりません")
		return
	}

	c.HTML(http.StatusOK, "edit", gin.H{"todo": todo})
}

// POST 作成
func Create(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.HTML(http.StatusBadRequest, "new", gin.H{
			"title": title,
			"error": "タイトルは必須です",
		})
		return
	}

	if len([]rune(title)) > 50 {
		c.HTML(http.StatusBadRequest, "new", gin.H{
			"title": title,
			"error": "タイトルは50文字以内で入力してください",
		})
		return
	}

	err := models.CreateTodo(title)
	if err != nil {
		c.String(http.StatusInternalServerError, "ToDoの作成に失敗しました")
		return
	}

	c.Redirect(http.StatusFound, "/")
}

// POST 変更
func Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "IDが不正です")
		return
	}

	title := c.PostForm("title")
	done := c.PostForm("done") == "on"

	todo, _ := models.GetTodoByID(id)

	if title == "" {
		todo.Title = title
		todo.Done = done

		c.HTML(http.StatusOK, "edit", gin.H{
			"todo":  todo,
			"error": "タイトルは必須です",
		})
		return
	}

	if len([]rune(title)) > 50 {
		todo.Title = title
		todo.Done = done

		fmt.Println("------------------------")
		fmt.Println(len(title))
		fmt.Println("------------------------")

		c.HTML(http.StatusOK, "edit", gin.H{
			"todo":  todo,
			"error": "タイトルは50文字以内で入力してください",
		})
		return
	}

	err = models.UpdateTodo(id, title, done)
	if err != nil {
		c.String(http.StatusInternalServerError, "ToDoの更新に失敗しました")
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "IDが不正です")
		return
	}

	err = models.DeleteTodo(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "ToDoの削除に失敗しました")
		return
	}

	c.Redirect(http.StatusFound, "/")
}
