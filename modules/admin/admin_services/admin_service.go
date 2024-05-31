package service

import (
	"github.com/natersland/akira-website-api-v2/data/repositories"
	authdto "github.com/natersland/akira-website-api-v2/modules/auth/auth_dto"
	authresponses "github.com/natersland/akira-website-api-v2/modules/auth/auth_responses"
)

type AdminService interface {
	CreateAdmin(req *authdto.CreateAdminDTO) (*authresponses.CreateAdminResponse, error)
}

type AdminServiceImpl struct {
	adminRepository repositories.AdminRepository
}

func NewAdminServiceImpl(adminRepository repositories.AdminRepository) AdminService {
	return &AdminServiceImpl{
		adminRepository: adminRepository,
	}
}

// CreateAdmin implements AdminService.
func (a *AdminServiceImpl) CreateAdmin(req *authdto.CreateAdminDTO) (*authresponses.CreateAdminResponse, error) {
	panic("unimplemented")
}
