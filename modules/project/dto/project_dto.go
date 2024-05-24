package dto

import "time"

type ProjectData struct {
	ProjectName  string    `json:"project_name" validate:"required"`
	ProjectUrl   string    `json:"project_url"`
	GitHubUrl    string    `json:"github_url"`
	ImageUrl     string    `json:"image_url"`
	Description  string    `json:"description"`
	Technologies []string  `json:"technologies"`
	Category     []string  `json:"category"`
	LauncedDate  time.Time `json:"launced_date"`
}
type CreateProjectDTO struct {
	ProjectData `json:"project"`
}

type UpdateProjectDTO struct {
	ID          string `json:"id" validate:"required"`
	ProjectData `json:"project"`
}

type DeleteProjectDTO struct {
	ID string `json:"id" validate:"required"`
}

type GetAllProjectDTO struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
