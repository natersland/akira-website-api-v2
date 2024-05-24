package entities

type Resume struct {
	FirstName   string   `gorm:"type:varchar(64);not null;column:first_name"`
	LastName    string   `gorm:"type:varchar(64);not null;column:last_name"`
	JobTitle    []string `gorm:"column:title" json:"title"`
	Email       string   `gorm:"column:email;not null"`
	ResumeUrl   string   `gorm:"type:varchar(256);column:resume_url"`
	LinkedIn    string   `gorm:"type:varchar(256);column:linkedin"`
	GitHub      string   `gorm:"type:varchar(256);column:github"`
	LastUpdated string   `gorm:"column:last_updated;not null"` // stamp this all update for show in website
}
