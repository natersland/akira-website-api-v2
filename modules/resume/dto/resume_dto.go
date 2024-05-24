package dto

type ResumeDataDto struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	JobTitle  []string `json:"title"`
	Email     string   `json:"email"`
	ResumeUrl string   `json:"resume_url"`
	LinkedIn  string   `json:"linkedin"`
	GitHub    string   `json:"github"`
}

type CreateResumeDto struct {
	ResumeDataDto `json:"data"`
}

type UpdateResumeDto struct {
	ResumeDataDto `json:"data"`
}
