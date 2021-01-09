package db

import (
	"../../clients/cassandra"
	"../../domain/access_token"
	"../../utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

// dbRepository struct & service has the same interface
func NewRepository() DbRepository {
	// Inorder to return this, it require dbRepository struct to have all the interface func
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
