package access_token

// import (
// 	"strings"

// 	"../../repository/db"
// 	"../../repository/rest"
// 	"../../utils/errors"
// )

// type Repository interface {
// 	GetById(string) (*AccessToken, *errors.RestErr)
// 	Create(AccessToken) *errors.RestErr
// 	UpdateExpirationTime(AccessToken) *errors.RestErr
// }

// type Service interface {
// 	GetById(string) (*AccessToken, *errors.RestErr)
// 	Create(AcessTokenRequest) (*AcessTokenRequest, *errors.RestErr)
// 	UpdateExpirationTime(AccessToken) *errors.RestErr
// }

// // service >> Service >> GetById
// type service struct {
// 	// repository   Repository
// 	restUserRepo rest.RestUsersRepository
// 	dbRepo           db.DbRepository
// }

// func NewService(repo Repository) Service {
// 	return &service{
// 		// repository: repo,
// 		restUsersRepo: usersRepo,
// 		dbRepoepo: dbRepo
// 	}
// }

// func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
// 	accessTokenId = strings.TrimSpace(accessTokenId)
// 	if len(accessTokenId) == 0 {
// 		return nil, errors.NewBadRequestError("invalid access token id")
// 	}
// 	accessToken, err := s.repository.GetById(accessTokenId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return accessToken, nil
// }

// func (s *service) Create(request AcessTokenRequest) (*AcessTokenRequest, *errors.RestErr) {
// 	if err := request.Validate(); err != nil {
// 		return nil, err
// 	}

// 	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	at := GetNewAccessToken(user.Id)
// 	at.Generate()

// 	if err := s.dbRepo.Create(at); err != nil {
// 		return nil, err
// 	}

// 	return &at, nil
// }

// func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
// 	if err := at.Validate(); err != nil {
// 		return err
// 	}

// 	return s.repository.UpdateExpirationTime(at)
// }
