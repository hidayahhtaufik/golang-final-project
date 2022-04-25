package router

import (
	"myGram/controllers"
	"myGram/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.DELETE("/", middlewares.Auth(), controllers.UserDelete)
		userRouter.PUT("/", middlewares.Auth(), controllers.UserUpdate)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Auth())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Auth())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComment)
		commentRouter.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socmedRouter := r.Group("/socialmedias")
	{
		socmedRouter.Use(middlewares.Auth())
		socmedRouter.POST("/", controllers.CreateSocmed)
		socmedRouter.GET("/", controllers.GetSocmed)
		socmedRouter.PUT("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocmed)
		socmedRouter.DELETE("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocmed)

	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running on port 8888")
	})

	return r
}
