package services

type TTESTService interface{}

type ttestService struct{}

func NewTTESTService() TTESTService {
	return &ttestService{}
}
