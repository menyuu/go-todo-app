package main

import (
	"go-todo-app/db"
	"go-todo-app/handlers"
	"go-todo-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	models.Migrate()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/static", "./static")

	// ルーティング
	r.GET("/", handlers.Index)
	r.GET("/:id", handlers.Show)
	r.GET("/new", handlers.New)
	r.GET("/edit/:id", handlers.Edit)
	r.POST("/create", handlers.Create)
	r.POST("/update/:id", handlers.Update)
	r.POST("/delete/:id", handlers.Delete)

	r.Run(":8080")
}
