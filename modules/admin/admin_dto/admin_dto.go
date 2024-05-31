package dto

type ChangePasswordDTO struct {
	AdminId     string `json:"admin_id" validate:"required"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type ChangeUsernameDTO struct {
	AdminId  string `json:"admin_id" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type ChangeEmailDTO struct {
	AdminId string `json:"admin_id" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

type DeleteAdminDTO struct {
	AdminId string `json:"admin_id" validate:"required"`
}
