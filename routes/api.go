package routes

import (
	"diluan/config"
	"diluan/controllers"
	"diluan/database"
	"diluan/middleware"
	"diluan/repositories"
	"diluan/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB        = database.InitMysql()
	mongodb *mongo.Database = config.InitMongoDB()

	//Repository
	testRepository               repositories.TestRepository                            = repositories.NewTestRepository(mongodb)
	userRepository               repositories.UserRepository                            = repositories.NewUserRepository(db)
	tokenRepository              repositories.OauthAccessTokenRepository                = repositories.NewOauthAccessTokenRepository(db)
	facebookPageRepository       repositories.FacebookPageRepository                    = repositories.NewFacebookPageRepository(db)
	fbPageConvoRepository        repositories.FacebookPageConversationRepository        = repositories.NewFacebookPageConversationRepository(db)
	fbPageConvoMessageRepository repositories.FacebookPageConversationMessageRepository = repositories.NewFacebookPageConversationMessageRepository(db, mongodb)
	//Service
	jwtService          services.JWTService          = services.NewJWTService()
	facebookPageService services.FacebookPageService = services.NewFacebookPageService(facebookPageRepository, fbPageConvoRepository, fbPageConvoMessageRepository)
	authService         services.AuthService         = services.NewAuthService(userRepository, tokenRepository)
	userService         services.UserService         = services.NewUserService(userRepository)
	queueService        services.QueueService        = services.NewQueueService()
	//Controller
	authController     controllers.AuthController     = controllers.NewAuthController(jwtService, authService, facebookPageService)
	userController     controllers.UserController     = controllers.NewUserController(jwtService, userService)
	facebookController controllers.FaceBookController = controllers.NewFaceBookController(facebookPageService, jwtService, queueService)
)

func InitRoutes(route *gin.RouterGroup) {
	authenRoutes(route)
	userRoutes(route)
	facebookRoutes(route)
}

func authenRoutes(route *gin.RouterGroup) {
	// route.POST("login", authController.Login)
	route.POST("register", authController.Register)

	route.GET("auth/facebook/callback", authController.FacebookAuthCallback)
	route.GET("auth/facebook/login", authController.FacebookAuthHandle)
}

func userRoutes(route *gin.RouterGroup) {
	userGroup := route.Group("user", middleware.AuthorizeJWT(jwtService))
	{
		userGroup.POST("info", userController.GetUserInfo)
	}
}

func facebookRoutes(route *gin.RouterGroup) {
	facebookGroup := route.Group("facebook", middleware.AuthorizeJWT(jwtService))
	{
		facebookGroup.GET("pages", facebookController.GetListPage)
		facebookGroup.GET("pages/:page_id/conversations", facebookController.GetPageConvo)
		facebookGroup.GET("pages/:page_id/conversations/:convo_id/messages", facebookController.GetConvoMessage)
	}
	// route.POST("facebook/pages/:id/conversations", facebookController.SyncConversation)

}
