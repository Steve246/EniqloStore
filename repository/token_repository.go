package repository

import (
	"encoding/hex"
	"eniqloStore/config"
	"eniqloStore/model/dto"
	"errors"
	"math/rand"

	"time"

	conn "github.com/spf13/viper"
	"gorm.io/gorm"
)

type TokenRepository interface {
	VerifyTokenV2(user_unique_id string) (bool, error)
	CreateTokenV2(user_unique_id string, length int) (string, error)
}

type tokenRepository struct {
	tokenConfig config.TokenConfig
	db          *gorm.DB
}

func (t *tokenRepository) CreateTokenV2(user_unique_id string, length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	token := hex.EncodeToString(b)
	expire := conn.GetInt64("token_auth.expire")
	expireDate := time.Now().Add(time.Second * time.Duration(expire))
	result := t.db.Exec("INSERT INTO authentication(user_unique_id,token_auth,expire) VALUES(?,?,?)", user_unique_id, token, expireDate)
	if result.Error != nil {
		return "", result.Error
	}
	return token, nil
}

func (t *tokenRepository) VerifyTokenV2(token_string string) (bool, error) {
	var userData dto.UserData

	result := t.db.Raw(`SELECT user_unique_id, expire FROM authentication WHERE token_auth = ? ORDER BY expire DESC LIMIT 1`, token_string).Scan(&userData)

	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("Unauthorized")
	}

	// check expire
	now := time.Now()
	date, _ := time.Parse(time.RFC3339, userData.Expire)
	if !now.Before(date) {
		return false, errors.New("Token Expired")
	}

	return true, nil
}

func NewTokenRepository(tokenConfig config.TokenConfig, dbClient *gorm.DB) TokenRepository {
	repo := new(tokenRepository)
	repo.tokenConfig = tokenConfig
	repo.db = dbClient
	return repo
}
