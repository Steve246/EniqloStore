package usecase

import (
	"eniqloStore/model/dto"
	"eniqloStore/repository"

	"github.com/golang-jwt/jwt"
)

// import (
// 	"7Zero4/model/dto"
// 	"7Zero4/repository"
// 	"log"

// 	"strconv"

// 	"github.com/golang-jwt/jwt"
// )

type TokenUsecase interface {
	VerifyAccessToken(tokenString string) (bool, error)
}

type tokenUsecase struct {
	tokenRepo repository.TokenRepository
}

type MyClaims struct {
	jwt.StandardClaims
	AuthToken dto.AuthToken
}

func (t *tokenUsecase) VerifyAccessToken(tokenString string) (bool, error) {
	return t.tokenRepo.VerifyTokenV2(tokenString)
}

func NewTokenUsecase(tokenRepo repository.TokenRepository) TokenUsecase {
	usecase := new(tokenUsecase)
	usecase.tokenRepo = tokenRepo
	return usecase
}
