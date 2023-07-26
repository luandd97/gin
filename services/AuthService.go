package services

import (
	"diluan/dto"
	"diluan/entities"
	"diluan/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	Create(user dto.RegisterDTO) entities.User
	StoreToken(token dto.FacebookTokenDTO) entities.OauthAccessToken
	IsDuplicateEmail(email string) bool
	FindByEmailAndDriver(email string, driver string) entities.User
	FindByEmail(email string) entities.User
	UpdateOrCreate(userID uint64, fbTokenDto dto.FacebookTokenDTO) entities.OauthAccessToken
}

type authService struct {
	userRepository  repositories.UserRepository
	tokenRepository repositories.OauthAccessTokenRepository
}

func NewAuthService(userReposity repositories.UserRepository, tokenRepository repositories.OauthAccessTokenRepository) AuthService {
	return &authService{
		userRepository:  userReposity,
		tokenRepository: tokenRepository,
	}
}

func (s *authService) IsDuplicateEmail(email string) bool {
	res := s.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (s *authService) Create(user dto.RegisterDTO) entities.User {
	userEntities := entities.User{}
	err := smapping.FillStruct(&userEntities, smapping.MapFields(user))

	if err != nil {
		log.Fatal("Failed Map %v", err)
	}

	res := s.userRepository.Create(userEntities)
	return res
}

func (s *authService) StoreToken(token dto.FacebookTokenDTO) entities.OauthAccessToken {
	oauthEntites := entities.OauthAccessToken{}
	err := smapping.FillStruct(&oauthEntites, smapping.MapFields(token))

	if err != nil {
		log.Fatal("Failed Map %v", err)
	}

	res := s.tokenRepository.Create(oauthEntites)
	return res
}

func (s *authService) FindByEmailAndDriver(email string, driver string) entities.User {
	return s.userRepository.FindByEmailAndDriver(email, driver)
}

func (s *authService) FindByEmail(email string) entities.User {
	return s.userRepository.FindByEmail(email)
}

func (s *authService) UpdateOrCreate(userID uint64, token dto.FacebookTokenDTO) entities.OauthAccessToken {

	oauthToken := s.tokenRepository.FindByUserID(userID)
	if oauthToken.ID == 0 {
		oauthEntites := entities.OauthAccessToken{}
		err := smapping.FillStruct(&oauthEntites, smapping.MapFields(token))
		if err != nil {
			log.Fatal("Failed Map %v", err)
		}
		createToken := s.tokenRepository.Create(oauthEntites)
		return createToken
	}
	oauthToken.AccessToken = token.AccessToken
	updateToken := s.tokenRepository.Update(oauthToken)
	return updateToken
}
