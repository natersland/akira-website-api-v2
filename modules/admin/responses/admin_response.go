package responses

import commonresponse "github.com/natersland/akira-website-api-v2/common/common_response"

type CreateAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"data"`
}

type ChangePasswordResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
}

type ChangeUsernameResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   struct {
		Id       string `json:"id"`
		Username string `json:"username"`
	}
}

type ChangeEmailResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   struct {
		Id    string `json:"id"`
		Email string `json:"email"`
	}
}

type GetAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
}
