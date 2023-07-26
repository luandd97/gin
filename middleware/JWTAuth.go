package middleware

import (
	"diluan/helpers"
	"diluan/services"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			response := helpers.BuildErrorResponse("Failed to process request", "Token not found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(header)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[issuer]", claims["issuer"])
		} else {
			response := helpers.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
