package handlers

import (
	"go-todo-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET 一覧
func TodoIndex(c *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		c.String(http.StatusInternalServerError, "ToDoの取得に失敗しました")
		return
	}

	c.HTML(http.StatusOK, "todos/index", gin.H{
		"todos": todos,
	})
}

func TodoShow(c *gin.Context) {
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

	c.HTML(http.StatusOK, "todos/show", gin.H{
		"todo": todo,
	})
}

// GET 作成
func TodoNew(c *gin.Context) {
	c.HTML(http.StatusOK, "todos/new", nil)
}

// GET 変更
func TodoEdit(c *gin.Context) {
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

	c.HTML(http.StatusOK, "todos/edit", gin.H{"todo": todo})
}

// POST 作成
func TodoCreate(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.HTML(http.StatusBadRequest, "new", gin.H{
			"title": title,
			"error": "タイトルは必須です",
		})
		return
	}

	if len([]rune(title)) > 50 {
		c.HTML(http.StatusBadRequest, "todos/new", gin.H{
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
func TodoUpdate(c *gin.Context) {
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

		c.HTML(http.StatusOK, "todos/edit", gin.H{
			"todo":  todo,
			"error": "タイトルは必須です",
		})
		return
	}

	if len([]rune(title)) > 50 {
		todo.Title = title
		todo.Done = done

		c.HTML(http.StatusOK, "todos/edit", gin.H{
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

func TodoDelete(c *gin.Context) {
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
