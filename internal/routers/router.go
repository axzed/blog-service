package routers

import "github.com/gin-gonic/gin"

// define router
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// define router group
	apiv1 := r.Group("/api/v1")
	{
		// handle tag
		apiv1.POST("/tags")
		apiv1.DELETE("/tags/:id")
		apiv1.PUT("/tags/:id")
		apiv1.PATCH("/tags/:id/state")
		apiv1.GET("/tags")

		// handle article
		apiv1.POST("/articles")
		apiv1.DELETE("/articles/:id")
		apiv1.PUT("/articles/:id")
		apiv1.PATCH("/articles/:id/state")
		apiv1.GET("/articles/:id")
		apiv1.GET("/articles")
	}

	return r
}