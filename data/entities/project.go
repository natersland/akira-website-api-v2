package entities

type Project struct {
	ID           string   `gorm:"column:id;primaryKey" json:"id"`
	ProjectName  string   `gorm:"column:project_name" json:"project_name"`
	ProjectUrl   string   `gorm:"column:project_url" json:"project_url"`
	GitHubUrl    string   `gorm:"column:github_url" json:"github_url"`
	ImageUrl     string   `gorm:"column:image_url" json:"image_url"`
	Description  string   `gorm:"column:description" json:"description"`
	Technologies []string `gorm:"column:technologies" json:"technologies"`
	Category     []string `gorm:"column:category" json:"category"`
	LauncedDate  string   `gorm:"column:launced_date" json:"launced_date"`
	LastUpdated  string   `gorm:"column:last_updated" json:"last_updated"` // stamp this all update for show in website
}
