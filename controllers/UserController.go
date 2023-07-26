package controllers

import (
	"diluan/helpers"
	"diluan/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserInfo(ctx *gin.Context)
}

type userController struct {
	jwtService  services.JWTService
	userService services.UserService
}

func NewUserController(jwtService services.JWTService, userService services.UserService) UserController {
	return &userController{
		jwtService:  jwtService,
		userService: userService,
	}
}

func (c *userController) GetUserInfo(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["UserID"])
	id, err := strconv.ParseUint(userID, 10, 64)

	if err != nil {
		resErr := helpers.BuildErrorResponse("User Not Found", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resErr)
		return
	}
	user := c.userService.FindById(id)
	res := helpers.BuildResponse(true, "OKE", user)
	ctx.JSON(http.StatusOK, res)

}
