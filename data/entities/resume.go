package entities

type Resume struct {
	FirstName   string   `gorm:"column:first_name" json:"first_name"`
	LastName    string   `gorm:"column:last_name" json:"last_name"`
	Title       []string `gorm:"column:title" json:"title"`
	Email       string   `gorm:"column:email" json:"email"`
	ResumeUrl   string   `gorm:"column:resume_url" json:"resume_url"`
	LinkedIn    string   `gorm:"column:linkedin" json:"linkedin"`
	GitHub      string   `gorm:"column:github" json:"github"`
	LastUpdated string   `gorm:"column:last_updated" json:"last_updated"` // stamp this all update for show in website
}
