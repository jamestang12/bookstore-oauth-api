package http

import (
	"net/http"

	atDomain "../domain/access_token"
	"../service/access_token"
	"../utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTojenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	//Implement access_token.Service so later on call call GetById in service to validate token
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTojenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := (c.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	// var at access_token.AccessToken
	// if err := c.ShouldBindJSON(&at); err != nil {
	// 	restErr := errors.NewBadRequestError("invalid json body")
	// 	c.JSON(restErr.Status, restErr)
	// 	return
	// }
	// if err := handler.service.Create(at); err != nil {
	// 	c.JSON(err.Status, err)
	// 	return
	// }
	// c.JSON(http.StatusCreated, at)
	var request atDomain.AcessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	access_token, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, access_token)

}
