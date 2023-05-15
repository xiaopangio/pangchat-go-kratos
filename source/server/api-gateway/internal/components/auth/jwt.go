package auth

import (
	"api-gateway/internal/conf"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtManager struct {
	secret []byte
}

func NewJwtManager(jwtCf *conf.Jwt) *JwtManager {
	secret := []byte(jwtCf.Key)
	return &JwtManager{secret: secret}
}

type UserToken struct {
	Uid          string `json:"uid"`
	NickName     string `json:"nick_name"`
	AccountId    string `json:"account_id"`
	PersonalDesc string `json:"personal_desc"`
	Avatar       string `json:"avatar"`
	City         string `json:"city"`
	Province     string `json:"province"`
	jwt.StandardClaims
}

func (u *JwtManager) SignUser(token *UserToken) (string, error) {
	now := time.Now()
	expire := now.Add(24 * time.Hour * 15)
	token.StandardClaims.ExpiresAt = expire.Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString(u.secret)
}
func (u *JwtManager) ParseUser(sign string) (*UserToken, error) {
	tokenClaims, err := jwt.ParseWithClaims(sign, &UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return u.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims == nil {
		return nil, errors.New("claims error")
	}
	if claims, ok := tokenClaims.Claims.(*UserToken); ok {
		return claims, nil
	}
	return nil, errors.New("claims error")
}

type SmsCodeToken struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func (u *JwtManager) SignSmsCode(phone string) (string, error) {
	now := time.Now()
	expire := now.Add(24 * time.Hour * 15)
	token := SmsCodeToken{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString(u.secret)
}
func (u *JwtManager) ParseSmsCode(sign string) (*SmsCodeToken, error) {
	tokenClaims, err := jwt.ParseWithClaims(sign, &SmsCodeToken{}, func(token *jwt.Token) (interface{}, error) {
		return u.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims == nil {
		return nil, errors.New("claims error")
	}
	if claims, ok := tokenClaims.Claims.(*SmsCodeToken); ok {
		return claims, nil
	}
	return nil, errors.New("claims error")
}
