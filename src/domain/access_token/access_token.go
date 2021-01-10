package access_token

import (
	"strings"
	"time"

	"../../utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
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

// Web - ClienT-Id: 123
// Mobile App - Clinet-Id: 234
