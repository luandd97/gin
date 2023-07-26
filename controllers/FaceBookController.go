package controllers

import (
	"diluan/entities"
	"diluan/helpers"
	"diluan/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type FaceBookController interface {
	GetListPage(ctx *gin.Context)
	SyncConversation(page entities.FacebookPage, userID uint64)
	GetConvoMessage(ctx *gin.Context)
	GetPageConvo(ctx *gin.Context)
}

type facebookController struct {
	fbPageService services.FacebookPageService
	jwtService    services.JWTService
	queueService  services.QueueService
}

func NewFaceBookController(
	fbPageService services.FacebookPageService,
	jwtService services.JWTService,
	queueService services.QueueService,
) FaceBookController {
	return &facebookController{
		fbPageService: fbPageService,
		jwtService:    jwtService,
		queueService:  queueService,
	}
}

func (c *facebookController) GetListPage(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["UserID"])
	id, err := strconv.ParseUint(userID, 10, 64)

	var facebookPages []entities.FacebookPage = c.fbPageService.GetByUserId(id)
	for _, page := range facebookPages {
		fmt.Println(page.PageID)
		go c.SyncConversation(page, id)
	}
	res := helpers.BuildResponse(true, "Oke", facebookPages)
	ctx.JSON(http.StatusOK, res)
}

func (c *facebookController) SyncConversation(page entities.FacebookPage, userID uint64) {
	if page.IsSyncConversation == true {
		return
	}
	page.IsSyncConversation = true
	c.fbPageService.Update(page)

	convo := c.fbPageService.GetPageConversation(page.AccessToken, page.PageID, userID)

	convoByte, _ := json.Marshal(convo)
	headers := make(map[string]interface{})
	headers["Sync_Conversation_From_Page"] = "convo"
	go c.queueService.MakePublish("LUMIA_QUEUE", convoByte, headers)
	return
}

func (c *facebookController) GetConvoMessage(ctx *gin.Context) {
	pageID := ctx.Param("page_id")
	convoID := ctx.Param("convo_id")
	fmt.Println(pageID, convoID)

	page := c.fbPageService.FindByPageID(pageID)
	if page.ID == 0 {
		res := helpers.BuildErrorResponse("Page Not Found", "Page Not Found", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	convo := c.fbPageService.FindByPageIDAndConvoID(pageID, convoID)
	if convo.ID == 0 {
		res := helpers.BuildErrorResponse("Conversation Not Found", "Conversation Not Found", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	messages := c.fbPageService.GetPageConversationDetail(page, convo)
	res := helpers.BuildResponse(true, "Succes", messages)
	ctx.JSON(http.StatusOK, res)
}

func (c *facebookController) GetPageConvo(ctx *gin.Context) {
	pageID := ctx.Param("page_id")
	userID := c.jwtService.GetUserID(ctx)
	page := c.fbPageService.FindByPageID(pageID)
	if page.ID == 0 {
		res := helpers.BuildErrorResponse("Page Not Found", "Page Not Found", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	convos := c.fbPageService.GetByPageIDAndUserID(pageID, userID)
	res := helpers.BuildResponse(true, "Success", convos)
	ctx.JSON(http.StatusOK, res)
}
