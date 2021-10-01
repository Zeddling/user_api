package routes

import (
	"net/http"

	"github.com/Zeddling/user/controllers/users"
	"github.com/gin-gonic/gin"
)

func StartGin(port int) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.DELETE("/user/:id", users.DeleteById)
		api.GET("/users", users.FindAllUsers)
		api.GET("/user/:id", users.FindUserByID)
		api.POST("/user", users.SaveUser)
		api.PUT("/user/:id", users.UpdateUser)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	return router
}
