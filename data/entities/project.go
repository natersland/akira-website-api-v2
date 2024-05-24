package entities

type Project struct {
	ID           string   `gorm:"column:id;primaryKey"`
	ProjectName  string   `gorm:"type:varchar(256);not null;column:project_name"`
	ProjectUrl   string   `gorm:"type:varchar(256);column:project_url"`
	GitHubUrl    string   `gorm:"type:varchar(256);column:github_url"`
	ImageUrl     string   `gorm:"type:varchar(256);column:image_url"`
	Description  string   `gorm:"type:varchar(256);column:description"`
	Technologies []string `gorm:"type:varchar(256);column:technologies"`
	Category     []string `gorm:"type:varchar(256);column:category"`
	LauncedDate  string   `gorm:"column:launced_date"`
	LastUpdated  string   `gorm:"column:last_updated;not null"` // stamp this all update for show in website
}
