package handlers

import (
	"go-todo-app/helpers"
	"go-todo-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET ユーザ一覧
func UserIndex(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, "ユーザー一覧の取得に失敗しました")
		return
	}

	currentUser, exists := c.Get("currentUser")
	if !exists {
		helpers.RenderHTML(c, http.StatusOK, "users/index", gin.H{
			"users": users,
		})
	}

	helpers.RenderHTML(c, http.StatusOK, "users/index", gin.H{
		"users":       users,
		"currentUser": currentUser,
	})
}

// GET マイページ
func UserShow(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, err := models.GetUserByID(id)
	if err != nil {
		c.String(http.StatusNotFound, "ユーザーが見つかりません")
		return
	}

	helpers.RenderHTML(c, http.StatusOK, "users/show", gin.H{
		"user": user,
	})
}

// GET ユーザ更新画面
func UserEdit(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, err := models.GetUserByID(id)
	if err != nil {
		helpers.RenderHTML(c, http.StatusNotFound, "user/edit", gin.H{
			"error": "ユーザーが見つかりません",
			"user":  user,
		})
		return
	}

	helpers.RenderHTML(c, http.StatusOK, "users/edit", gin.H{
		"user": user,
	})
}

// POST 変更
func UserUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUserByID(id)

	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	if c.PostForm("password") != "" {
		user.Password = c.PostForm("password")
	}

	err := models.UpdateUser(user)
	if err != nil {
		helpers.RenderHTML(c, http.StatusBadRequest, "user/edit", gin.H{
			"error": "入力情報が正しくありません",
			"user":  user,
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/users/"+idStr)
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")
	err := models.DeleteUser(id)

	if err != nil {
		c.String(http.StatusInternalServerError, "ユーザーの削除に失敗しました")
	}

	c.Redirect(http.StatusSeeOther, "/users")
}
