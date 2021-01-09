package app

import (
	"../clients/cassandra"
	"../domain/access_token"
	"../http"
	"../repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	//Repository & DbRepository has the same interface
	dbRepository := db.NewRepository()
	//A new service need to work with a dbRepository
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8000")

}
