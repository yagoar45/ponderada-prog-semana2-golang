// handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetUserHandler(c *gin.Context) {
	id := c.Param("id")

	user := User{
		ID:   1,
		Name: "John Doe",
	}

	if id != "1" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
