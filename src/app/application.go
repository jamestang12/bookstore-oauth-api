package app

import (
	"../domain/access_token"
	"../http"
	"../repository/db"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// session, dbErr := cassandra.GetSession()
	// if dbErr != nil {
	// 	panic(dbErr)
	// }
	// session.Close()

	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "oauth"

	var session *gocql.Session
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
	session.Close()

	//Repository & DbRepository has the same interface
	dbRepository := db.NewRepository()
	//A new service need to work with a dbRepository
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8000")

}
