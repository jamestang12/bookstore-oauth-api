package http

import (
	"net/http"

	"../domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTojenHandler interface {
	GetById(*gin.Context)
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
