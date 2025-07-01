package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RenderHTML(c *gin.Context, code int, templateName string, data gin.H) {
	user, _ := c.Get("currentUser")
	if data == nil {
		data = gin.H{}
	}
	data["currentUser"] = user

	log.Println("-----------------------")
	log.Println(data)
	log.Println("-----------------------")

	c.HTML(code, templateName, data)
}
