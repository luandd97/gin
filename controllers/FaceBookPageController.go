package controllers

type FaceBookPageController interface{}

type facebookpageController struct{}

func NewFaceBookPageController() FaceBookPageController {
	return &facebookpageController{}
}
