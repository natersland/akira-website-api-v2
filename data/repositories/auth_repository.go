package repositories

import (
	commonhelpers "github.com/natersland/akira-website-api-v2/common/common_helpers"
	commonresponse "github.com/natersland/akira-website-api-v2/common/common_response"
	"github.com/natersland/akira-website-api-v2/data/entities"
	authdto "github.com/natersland/akira-website-api-v2/modules/auth/auth_dto"
	authhelpers "github.com/natersland/akira-website-api-v2/modules/auth/auth_helpers"
	authresponses "github.com/natersland/akira-website-api-v2/modules/auth/auth_responses"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateAdmin(req *authdto.CreateAdminDTO) (*authresponses.CreateAdminResponse, error)
}

type AuthRepositoryImpl struct {
	*gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

// CreateAdmin implements AuthRepository.
func (a *AuthRepositoryImpl) CreateAdmin(req *authdto.CreateAdminDTO) (*authresponses.CreateAdminResponse, error) {

	hashedPassword, err := authhelpers.BcryptHashingPassword(req.Password)
	if err != nil {
		return &authresponses.CreateAdminResponse{
			Status: commonresponse.CommonResponse{
				Success: false,
				Message: "Repo - Failed to hash password: " + err.Error(),
			},
			Data: nil,
		}, err
	}

	// Save to database
	admin := entities.Admin{
		Id:        commonhelpers.GenerateUUID(),
		UserName:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		UpdatedAt: commonhelpers.GetCurrentTimeISO(),
	}

	// Save to database
	result := a.DB.Create(&admin)
	if result.Error != nil {
		return &authresponses.CreateAdminResponse{
			Status: commonresponse.CommonResponse{
				Success: false,
				Message: "Repo - Failed to create admin: " + result.Error.Error(),
			},
			Data: nil,
		}, result.Error
	}

	return &authresponses.CreateAdminResponse{
		Status: commonresponse.CommonResponse{
			Success: true,
			Message: "Repo - Successfully created admin",
		},
		Data: &authresponses.AdminFieldResponse{
			Id:       admin.Id,
			Username: admin.UserName,
			Email:    admin.Email,
		},
	}, nil

}
