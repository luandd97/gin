package controllers

type FacebookPageConversationController interface{}

type facebookpageconversationController struct{}

func NewFacebookPageConversationController() FacebookPageConversationController {
	return &facebookpageconversationController{}
}
