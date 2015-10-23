package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var pessoas []map[string]interface{} = make([]map[string]interface{}, 0)

func main() {
	r := gin.Default()

	r.GET("/pessoas", func(c *gin.Context) {
		c.JSON(200, pessoas)
	})
	r.POST("/pessoas", func(c *gin.Context) {
		pessoa := make(map[string]interface{})
		c.BindJSON(&pessoa)

		pessoas = append(pessoas, pessoa)

		c.AbortWithStatus(http.StatusOK)
	})
	r.DELETE("/pessoas/:name", func(c *gin.Context) {
		name := c.Param("name")

		newCollection := make([]map[string]interface{}, 0)

		for _, obj := range pessoas {
			if !strings.EqualFold(name, obj["name"].(string)) {
				newCollection = append(newCollection, obj)
			}
		}

		pessoas = newCollection

		c.AbortWithStatus(http.StatusOK)
	})

	r.Run(":8080")
}