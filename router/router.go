package router

import (
	"compartamos_go/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UsersController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Bienvenido....")
	})

	baseRouter := router.Group("/api")
	userRouter := baseRouter.Group("/users")
	userRouter.GET("", userController.FindAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.POST("", userController.Create)
	userRouter.PATCH("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)
	return router
}
