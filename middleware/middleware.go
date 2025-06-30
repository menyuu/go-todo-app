package middleware

import (
	"go-todo-app/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		idFloat, ok := userID.(uint)
		if !ok {
			c.Set("currentUser", nil)
			c.Next()
			return
		}

		id := int(idFloat)
		user, err := models.GetUserByID(id)
		if err != nil {
			c.Set("currentUser", nil)
			c.Next()
			return
		}

		c.Set("currentUser", &user)
		c.Next()
	}
}
