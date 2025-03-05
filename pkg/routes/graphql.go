package routes

import (
	"aas.dev/pkg/graph"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func SetupGraphRoutes(router *gin.Engine) {
	router.POST("/graphql", func(c *gin.Context) {
		var request map[string]interface{}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		params := graphql.Params{
			Schema:        graph.Schema,
			RequestString: request["query"].(string),
		}
		result := graphql.Do(params)
		if len(result.Errors) > 0 {
			c.JSON(400, result.Errors)
		} else {
			c.JSON(200, result.Data)
		}
	})
}
