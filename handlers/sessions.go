package handlers

import (
	"go-todo-app/forms"
	"go-todo-app/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowSignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "sessions/new", nil)
}

func SignUp(c *gin.Context) {
	var form forms.SignUpForm

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "sessions/new", gin.H{
			"error": "入力を確認してください",
		})
		return
	}

	errors := forms.ValidateStruct(form)
	if len(errors) > 0 {
		c.HTML(http.StatusBadRequest, "sessions/new", gin.H{
			"errors": errors,
			"form":   form,
		})
		return
	}

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := models.CreateUser(&user); err != nil {
		c.HTML(http.StatusInternalServerError, "sessions/new", gin.H{
			"error": "登録に失敗しました",
			"form":  form,
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusSeeOther, "/")
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "sessions/login", nil)
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := models.AuthenticateUser(email, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "users/login", gin.H{
			"error": "ユーザー認証に失敗しました",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.Redirect(http.StatusSeeOther, "/")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusSeeOther, "/login")
}
