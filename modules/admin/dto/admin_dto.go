package dto

type CreateAdminDTO struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ChangePasswordDTO struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type ChangeUsernameDTO struct {
	Username string `json:"username" validate:"required"`
}

type ChangeEmailDTO struct {
	Email string `json:"email" validate:"required,email"`
}
