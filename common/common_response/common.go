package commonresponse

type CommonResponse struct {
	Success bool   `gorm:"column:is_success" json:"is_success"`
	Message string `gorm:"column:message" json:"message"`
}
