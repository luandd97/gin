package controllers

import (
	"diluan/config"
	"diluan/dto"
	"diluan/entities"
	"diluan/helpers"
	"diluan/services"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var (
	MessageErrorRequest = "Failed To Process Request"
)

type AuthController interface {
	// Login(ctx *gin.Context, driver string)
	Register(ctx *gin.Context)
	FacebookAuthHandle(ctx *gin.Context)
	FacebookAuthCallback(ctx *gin.Context)
}

type authController struct {
	jwtService          services.JWTService
	authService         services.AuthService
	facebookPageService services.FacebookPageService
}

type FacebookUserDetails struct {
	ID    string
	Name  string
	Email string
	Phone string
}

func NewAuthController(jwtService services.JWTService, authService services.AuthService, facebookPageService services.FacebookPageService) AuthController {
	return &authController{
		jwtService:          jwtService,
		authService:         authService,
		facebookPageService: facebookPageService,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO

	err := ctx.ShouldBind(&registerDTO)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed To Process Request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helpers.BuildErrorResponse("Failed To Process Request", "Duplicate Emaiul", nil)
		ctx.JSON(http.StatusConflict, response)
	} else {
		create := c.authService.Create(registerDTO)
		token := ""
		create.AccessToken = token
		res := helpers.BuildResponse(true, "Success", create)
		ctx.JSON(http.StatusCreated, res)
	}

}

func (c *authController) FacebookAuthHandle(ctx *gin.Context) {
	conf := config.GetFacebookOAuthConfig()
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	res := helpers.BuildResponse(true, "Success", url)
	ctx.JSON(http.StatusOK, res)

}

func (c *authController) FacebookAuthCallback(ctx *gin.Context) {
	var code = ctx.Query("code")
	var OAuth2Config = config.GetFacebookOAuthConfig()

	token, err := OAuth2Config.Exchange(oauth2.NoContext, code)
	if err != nil || token == nil {
		res := helpers.BuildErrorResponse("Failed Login Facebook", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	fbUserDetails, fbUserDetailsError := GetUserInfoFromFacebook(token.AccessToken)
	if fbUserDetailsError != nil {
		res := helpers.BuildErrorResponse("Get User Failed", fbUserDetailsError.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	fbUserDetails.Email = "duongdiluan@yahoo.com"
	authToken, authTokenError := c.SignInUser(fbUserDetails)
	if authTokenError != nil {
		res := helpers.BuildErrorResponse("Login Failed", authTokenError.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}

	access_token := authToken.AccessToken
	fbTokenDto := dto.FacebookTokenDTO{
		Name:           "Facebook",
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		ExpiresAt:      token.Expiry.AddDate(0, 0, 7),
		UserID:         authToken.ID,
		FacebookUserId: fbUserDetails.ID,
	}
	c.authService.UpdateOrCreate(authToken.ID, fbTokenDto)
	go c.facebookPageService.HandleFacebookPage(token.AccessToken, fbUserDetails.ID, authToken.ID)

	ctx.Redirect(http.StatusTemporaryRedirect, os.Getenv("FE_APP_URL")+"/auth/facebook/callback?access_token="+access_token)
}

func GetUserInfoFromFacebook(token string) (FacebookUserDetails, error) {
	var fbUserDetails FacebookUserDetails
	facebookUserDetailsRequest, _ := http.NewRequest("GET", "https://graph.facebook.com/me?fields=id,name,email&access_token="+token, nil)
	facebookUserDetailsResponse, facebookUserDetailsResponseError := http.DefaultClient.Do(facebookUserDetailsRequest)
	if facebookUserDetailsResponseError != nil {
		return FacebookUserDetails{}, errors.New("Error occurred while getting information from Facebook")
	}
	decoder := json.NewDecoder(facebookUserDetailsResponse.Body)

	decoderErr := decoder.Decode(&fbUserDetails)

	defer facebookUserDetailsResponse.Body.Close()

	if decoderErr != nil {
		return FacebookUserDetails{}, errors.New("Error occurred while getting information from Facebook")
	}
	fbUserDetails.Email = "duongdiluan@yahoo.com"
	return fbUserDetails, nil
}

func (c *authController) SignInUser(facebookUserDetails FacebookUserDetails) (entities.User, error) {
	// var result UserDetails
	if facebookUserDetails == (FacebookUserDetails{}) {
		return entities.User{}, errors.New("User details Can't be empty")
	}

	if facebookUserDetails.Email == "" {
		return entities.User{}, errors.New("Email can't be empty")
	}

	if facebookUserDetails.Name == "" {
		return entities.User{}, errors.New("Name can't be empty")
	}

	registerDTO := dto.RegisterDTO{
		Email:        facebookUserDetails.Email,
		Name:         facebookUserDetails.Name,
		Phone:        facebookUserDetails.Phone,
		Password:     "",
		Status:       true,
		SocialDriver: "facebook",
	}
	user := c.authService.FindByEmailAndDriver(facebookUserDetails.Email, registerDTO.SocialDriver)
	if user.ID == 0 {
		create := c.authService.Create(registerDTO)
		token := c.jwtService.GenerateToken(uint64(create.ID))
		create.AccessToken = token
		return create, nil
	}
	token := c.jwtService.GenerateToken(uint64(user.ID))
	user.AccessToken = token
	return user, nil
}
