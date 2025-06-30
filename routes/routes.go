package routes

import (
	"go-todo-app/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	// todos関連
	r.GET("/", handlers.TodoIndex)
	r.GET("/todos/new", handlers.TodoNew)
	r.GET("/todos/edit/:id", handlers.TodoEdit)
	r.POST("/todos/create", handlers.TodoCreate)
	r.GET("/todos/:id", handlers.TodoShow)
	r.POST("/todos/update/:id", handlers.TodoUpdate)
	r.POST("/todos/delete/:id", handlers.TodoDelete)

	// ユーザー関連
	r.GET("/users", handlers.UserIndex)
	r.GET("/users/:id", handlers.UserShow)
	r.GET("/users/edit/:id", handlers.UserEdit)
	r.POST("users/update/:id", handlers.UserUpdate)
	r.POST("users/delete/:id", handlers.UserDelete)

	// ユーザー認証関連
	r.GET("/signup", handlers.ShowSignUp)
	r.POST("/signup", handlers.SignUp)
	r.GET("/login", handlers.ShowLogin)
	r.POST("/login", handlers.Login)
	r.POST("/logout", handlers.Logout)
}
