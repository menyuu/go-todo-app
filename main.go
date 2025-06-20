package main

import (
	"go-todo-app/db"
	"go-todo-app/handlers"
	"go-todo-app/models"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob("templates/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	partials, err := filepath.Glob("templates/partials/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob("templates/**/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		partialCopy := make([]string, len(partials))
		copy(partialCopy, partials)
		files := append(layoutCopy, partialCopy...)
		files = append(files, include)
		name := strings.TrimSuffix(filepath.Base(include), filepath.Ext(include))
		r.AddFromFiles(name, files...)
	}

	return r
}

func main() {
	db.Connect()
	models.Migrate()

	r := gin.Default()
	r.HTMLRender = createMyRender()
	r.Static("/static", "./static")

	// ルーティング
	r.GET("/", handlers.Index)
	r.GET("/new", handlers.New)
	r.GET("/edit/:id", handlers.Edit)
	r.POST("/create", handlers.Create)
	r.GET("/:id", handlers.Show)
	r.POST("/update/:id", handlers.Update)
	r.POST("/delete/:id", handlers.Delete)

	r.Run(":8080")
}
