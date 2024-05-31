package responses

import commonresponse "github.com/natersland/akira-website-api-v2/common/common_response"

type AdminFieldResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetAdminsResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   []AdminFieldResponse          `json:"data"`
}

type GetAdminDetailResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   *AdminFieldResponse           `json:"data"`
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

type DeleteAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
}
