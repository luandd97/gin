package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTService interface {
	GenerateToken(userID uint64) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserID(ctx *gin.Context) uint64
}

type jwtService struct {
	secretKey string
	issuer    string
}

type jwtCustomeClaim struct {
	UserID uint64 `json:user_id`
	jwt.StandardClaims
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "NUHA%JK@#!K",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "NUHA%JK@#!K"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID uint64) string {
	curentTime := time.Now()
	clains := &jwtCustomeClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: curentTime.Add(time.Hour * 24).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) GetUserID(ctx *gin.Context) uint64 {
	authHeader := ctx.GetHeader("Authorization")
	token, err := j.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["UserID"])
	id, _ := strconv.ParseUint(userID, 10, 64)
	return id
}
