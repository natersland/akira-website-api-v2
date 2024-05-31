package authresponses

import commonresponse "github.com/natersland/akira-website-api-v2/common/common_response"

type AdminFieldResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   *AdminFieldResponse           `json:"data"`
}
