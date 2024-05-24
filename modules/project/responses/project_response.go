package responses

import (
	commonresponse "github.com/natersland/akira-website-api-v2/common/common_response"
	"github.com/natersland/akira-website-api-v2/modules/project/dto"
)

type CreateProjectResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   dto.ProjectData               `json:"data"`
}

type UpdateProjectResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   dto.ProjectData               `json:"data"`
}

type DeleteProjectResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
}

type GetAllProjectsResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   []dto.ProjectData             `json:"data"`
}
