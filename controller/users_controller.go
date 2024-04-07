package controller

import (
	"compartamos_go/data/request"
	"compartamos_go/data/response"
	"compartamos_go/helper"
	"compartamos_go/service"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors" // Importa el paquete CORS de Gin
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{
		usersService: service,
	}
}

// Configura las opciones CORS
var corsConfig = cors.DefaultConfig()

func init() {
	// Permite solicitudes desde cualquier origen
	corsConfig.AllowAllOrigins = true
}

func (controller *UsersController) Create(ctx *gin.Context) {
	createUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.ErrorSync(err)

	controller.usersService.Create(createUsersRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UsersController) Update(ctx *gin.Context) {
	updateUsersRequest := request.UpdateUsersRequest{}
	err := ctx.ShouldBindJSON(&updateUsersRequest)
	helper.ErrorSync(err)

	userId := ctx.Param("userId")

	id, err := strconv.Atoi(userId)

	helper.ErrorSync(err)
	updateUsersRequest.Id = id

	controller.usersService.Update(updateUsersRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UsersController) Delete(ctx *gin.Context) {

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorSync(err)

	controller.usersService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UsersController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorSync(err)

	userResponse := controller.usersService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UsersController) FindAll(ctx *gin.Context) {
	userResponse := controller.usersService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
