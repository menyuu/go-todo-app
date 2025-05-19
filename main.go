package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo は1つのタスクを表す構造体
type Todo struct {
	ID    int
	Title string
	Done  bool
}

// todos はメモリ上のToDoリスト（仮のDB）
var todos []Todo
var nextID = 1

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"todos": todos})
	})

	r.POST("/todos", func(c *gin.Context) {
		title := c.PostForm("title")
		if title == "" {
			c.String(http.StatusBadRequest, "タイトルが必要です")
			return
		}

		newTodo := Todo{
			ID:    nextID,
			Title: title,
			Done:  false,
		}
		nextID++
		todos = append(todos, newTodo)

		c.Redirect(http.StatusFound, "/")
	})

	// // GET /todos → 一覧取得
	// r.GET("/todos", func(c *gin.Context) {
	// 	if todos == nil {
	// 		c.JSON(http.StatusOK, []Todo{})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, todos)
	// })

	// // POST /todos → 新規作成
	// r.POST("/todos", func(c *gin.Context) {
	// 	var newTodo Todo
	// 	if err := c.ShouldBindJSON(&newTodo); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	newTodo.ID = nextID
	// 	nextID++
	// 	todos = append(todos, newTodo)

	// 	c.JSON(http.StatusCreated, newTodo)
	// })

	// r.StaticFile("/", "./frontend/index.html")

	r.Run(":8080")
}
