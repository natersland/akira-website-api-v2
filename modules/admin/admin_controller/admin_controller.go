package controller

import service "github.com/natersland/akira-website-api-v2/modules/admin/admin_services"

type AdminController interface {
}

type AdminControllerImpl struct {
	adminService service.AdminService
}

func NewAdminControllerImpl(adminService service.AdminService) AdminController {
	return &AdminControllerImpl{
		adminService: adminService,
	}
}
