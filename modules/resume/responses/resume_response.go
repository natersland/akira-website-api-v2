package responses

import commonresponse "github.com/natersland/akira-website-api-v2/common/common_response"

type GetResumeResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   struct {
		FullName    string   `json:"full_name"`
		JobTitle    []string `json:"job_title"`
		Email       string   `json:"email"`
		ResumeUrl   string   `json:"resume_url"`
		LinkedIn    string   `json:"linkedin"`
		GitHub      string   `json:"github"`
		LastUpdated string   `json:"last_updated"`
	} `json:"data"`
}
