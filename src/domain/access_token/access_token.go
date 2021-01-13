package access_token

import (
	"fmt"
	"strings"
	"time"

	"../../utils/crypto_utils"
	"../../utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentails = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AcessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	// Used for password grant_type
	Password string `json:"password"`
	Username string `json:"username"`

	// Used for client_credentials
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AcessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentails:
		break
	default:
		return errors.NewBadRequestError("invalid grant type")
	}

	return nil
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	// if len(at.AccessToken) == 0 {
	// 	return errors.NewBadRequestError("invalid access token id")
	// }
	if at.AccessToken == "" {
		// return errors.NewBadRequestError("invalid access token id")
		return errors.NewBadRequestError("invaild access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invaild access user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invaild access client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invaild expiration time")
	}

	return nil
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}

// Web - ClienT-Id: 123
// Mobile App - Clinet-Id: 234
