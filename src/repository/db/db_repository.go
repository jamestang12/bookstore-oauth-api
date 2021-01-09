package db

import (
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

func (*dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
