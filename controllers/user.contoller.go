package controllers

import (
	"net/http"

	"example.com/gin-api/models"
	"example.com/gin-api/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Userservices services.Userservices
}

func New(userservices services.Userservices) UserController {
	return UserController{
		Userservices: userservices,
	}
}

func (uc *UserController) Createuser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"mssage": err.Error()})
		return
	}
	if err := uc.Userservices.Createuser(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"mssage": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "sucesss"})
}

func (uc *UserController) Getuser(ctx *gin.Context) {
	username := ctx.Param("name")
	data, err := uc.Userservices.Getuser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"mssage": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.Userservices.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"mssage": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})

}

func (uc *UserController) Updateuser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"mssage": err.Error()})
		return
	}
	if err := uc.Userservices.Updateuser(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"mssage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "sucesss"})
}

func (uc *UserController) Deleteuser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.Userservices.Deleteuser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"mssage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "sucesss"})

}

func (uc *UserController) RegisterUserRouter(rg *gin.RouterGroup) {
	userrouter := rg.Group("/user")
	userrouter.POST("/create", uc.Createuser)
	userrouter.GET("/get/:name", uc.Getuser)
	userrouter.GET("/getall", uc.GetAll)
	userrouter.PATCH("/update", uc.Updateuser)
	userrouter.DELETE("/delete:name", uc.Deleteuser)

}
