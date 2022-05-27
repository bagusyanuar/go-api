package controller

import (
	"errors"
	"go-api/src/lib"
	"go-api/src/model"
	"go-api/src/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Subject(c *gin.Context) {
	if c.Request.Method == "POST" {
		name := c.PostForm("name")

		subject := model.Subject{
			Name: strings.Title(name),
			Slug: lib.MakeSlug(name),
			Icon: nil,
		}

		err := repository.CreateSubject(&subject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"data":    nil,
				"message": "failed to insert " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    nil,
			"message": "success",
		})
		return
	}

	var subjects []model.Subject
	q := c.Query("q")
	data, err := repository.FindSubject(&subjects, q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "failed to fetch data " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": "success",
	})
}

func SubjectByID(c *gin.Context) {
	id := c.Param("id")

	var subject model.Subject
	data, err := repository.FindSubjectByID(&subject, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"data":    nil,
				"message": "success",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    nil,
			"message": "failed to fetch detail subject " + err.Error(),
		})
		return
	}

	if c.Request.Method == "PATCH" {
		id := data.ID
		name := c.PostForm("name")

		data_update := map[string]interface{}{
			"name": strings.Title(name),
			"slug": lib.MakeSlug(name),
		}
		err = repository.PatchSubjectByID(id, data_update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"data":    nil,
				"message": "failed to pathc subject " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"data":    nil,
			"message": "success",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": "success",
	})
}
