package controller

import (
	"encoding/json"
	"go-api/database"
	"go-api/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func User(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")
	var users []model.User
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")
		roles, _ := json.Marshal([]string{"ROLE_MEMBER"})
		provider, _ := json.Marshal([]string{"app"})

		hashedPassword, errHashed := bcrypt.GenerateFromPassword([]byte(password), 13)
		if errHashed != nil {
			c.AbortWithError(http.StatusInternalServerError, errHashed)
		}
		var vPassword string = string(hashedPassword)
		user := model.User{
			Email:    email,
			Username: username,
			Password: &vPassword,
			Roles:    roles,
			Provider: provider,
		}

		if err := database.CONFIG.Create(&user).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   user,
		})
		return
	}
	if err := database.CONFIG.Find(&users).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   users,
	})
}
