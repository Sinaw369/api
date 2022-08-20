package controllers

import (
	"net/http"

	"book.com/sina-apis/models"
	"book.com/sina-apis/services"
	"github.com/gin-gonic/gin"
)

type AutherController struct {
	AutherService services.AutherService
}
func New(autherService  services.AutherService)AutherController{
	return AutherController{
		AutherService: autherService,
	}

}

func (ac *AutherController) AddBook(ctx *gin.Context) {
	var book models.Book
	if err:=ctx.ShouldBindJSON(&book);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"massage":err.Error()})
		return
	}
	err:=ac.AutherService.AddBook(&book)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"massage":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"massage":"success"})
}

func (ac *AutherController) GetAllAutherBook(ctx *gin.Context) {
	authername:= ctx.Param("name")
	books,err:= ac.AutherService.GetAllAutherBook(&authername)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"massage":err.Error()})
	}
	ctx.JSON(http.StatusOK,books)
}
func (ac *AutherController) RemoveBook(ctx *gin.Context){
	bn:=ctx.Param("bookname")
	an:=ctx.Param("authername")
	if err:=ac.AutherService.RemoveBook(&bn,&an);err!=nil{
		ctx.JSON(http.StatusBadGateway,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,"success")
}
func (ac *AutherController) UpdateBook(ctx *gin.Context){
	var book models.Book 
	if err:=ctx.ShouldBindJSON(&book);err !=nil{
		ctx.JSON(http.StatusBadRequest,err.Error())
		return
	}
	err:=ac.AutherService.UpdateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusBadGateway,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"massage":"success"})
}
func (ac *AutherController)RegisterAutherRouter(rg *gin.RouterGroup){
	autherroute:=rg.Group("/auther")
	autherroute.GET("/getall/:name",ac.GetAllAutherBook)
	autherroute.POST("/add",ac.AddBook)
	autherroute.PATCH("update",ac.UpdateBook)
	autherroute.DELETE("/del/:bookname/:authername",ac.RemoveBook)
}

