package handlers

import (
	"go-todo-app/forms"
	"go-todo-app/helpers"
	"go-todo-app/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GET 一覧
func TodoIndex(c *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		c.String(http.StatusInternalServerError, "ToDoの取得に失敗しました")
		return
	}

	helpers.RenderHTML(c, http.StatusOK, "todos/index", gin.H{
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

	helpers.RenderHTML(c, http.StatusOK, "todos/show", gin.H{
		"todo": todo,
	})
}

// GET 作成
func TodoNew(c *gin.Context) {
	helpers.RenderHTML(c, http.StatusOK, "todos/new", nil)
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

	helpers.RenderHTML(c, http.StatusOK, "todos/edit", gin.H{"todo": todo})
}

// POST 作成
func TodoCreate(c *gin.Context) {
	var form forms.TodoForm
	if err := c.ShouldBind(&form); err != nil {
		helpers.RenderHTML(c, http.StatusBadRequest, "todos/new", gin.H{
			"error": "入力エラーがあります",
		})
		return
	}

	errors := forms.ValidateStruct(form)
	if len(errors) > 0 {
		helpers.RenderHTML(c, http.StatusBadRequest, "todos/new", gin.H{
			"error": errors,
			"form":  form,
		})
	}

	session := sessions.Default(c)
	userID, ok := session.Get("user_id").(uint)
	if !ok {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	todo := models.Todo{
		Title:  form.Title,
		Done:   false,
		UserID: userID,
	}

	if err := models.CreateTodo(&todo); err != nil {
		helpers.RenderHTML(c, http.StatusInternalServerError, "todos/new", gin.H{
			"error": "保存に失敗しました",
			"form":  form,
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

// POST 変更
func TodoUpdate(c *gin.Context) {
	idStr := c.Param("id")

	var form forms.TodoForm
	if err := c.ShouldBind(&form); err != nil {
		helpers.RenderHTML(c, http.StatusBadRequest, "todos/edit", gin.H{
			"error": "入力エラーがあります",
		})
		return
	}

	id, _ := strconv.Atoi(idStr)

	errors := forms.ValidateStruct(form)
	if len(errors) > 0 {
		helpers.RenderHTML(c, http.StatusBadRequest, "todos/edit", gin.H{
			"error": errors,
			"form":  form,
			"id":    id,
		})
		return
	}

	session := sessions.Default(c)
	userID, ok := session.Get("user_id").(uint)
	if !ok {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	todo, err := models.GetTodoByID(id)
	if err != nil {
		helpers.RenderHTML(c, http.StatusNotFound, "todos/edit", gin.H{
			"error": "Todoが見つかりません",
			"form":  form,
			"id":    id,
		})
		return
	} else if todo.UserID != userID {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	todo.Title = form.Title
	todo.Done = form.Done

	if err := models.UpdateTodo(todo); err != nil {
		helpers.RenderHTML(c, http.StatusInternalServerError, "todos/edit", gin.H{
			"error": "更新に失敗しました",
			"form":  form,
			"id":    id,
		})
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
