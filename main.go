package main

import (
	"go-todo-app/db"
	"go-todo-app/middleware"
	"go-todo-app/models"
	"go-todo-app/routes"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
		relPath := strings.TrimPrefix(include, "templates/")
		name := strings.TrimSuffix(relPath, filepath.Ext(include))
		r.AddFromFiles(name, files...)
	}

	return r
}

func main() {
	db.Connect()
	models.Migrate()

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("todo_session", store))
	r.Use(middleware.CurrentUser())

	r.HTMLRender = createMyRender()
	r.Static("/static", "./static")
	// ルーティング
	routes.SetUpRoutes(r)

	r.Run(":8080")
}
