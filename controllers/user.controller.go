package controllers

import (
	"net/http"

	"book.com/sina-apis/services"
	"github.com/gin-gonic/gin"

)

type UserController struct {
	UserService services.UserService
}
func New1(userservice services.UserService)UserController{
	return UserController{
		UserService: userservice,
	}
}
func (uc *UserController)BuyBook(ctx *gin.Context){
     bn:=ctx.Param("bookname")
	 an:=ctx.Param("authername")
	 result,err:=uc.UserService.BuyBook(&bn,&an)
	 if err != nil {
		ctx.JSON(http.StatusBadGateway,err.Error())
	 }
	 ctx.JSON(http.StatusOK,result)
	
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup){
	userroute:=rg.Group("/user")
	userroute.GET("/buy/:bookname/:authername",uc.BuyBook)
}