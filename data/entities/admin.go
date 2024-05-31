package entities

type Admin struct {
	Id        string `gorm:"column:id;primaryKey"`
	UserName  string `gorm:"type:varchar(256);not null;column:username" validate:"required" json:"username"`
	Email     string `gorm:"type:varchar(256);not null;column:email" validate:"required" json:"email"`
	Password  string `gorm:"type:varchar(256);not null;column:password" validate:"required" json:"password"`
	UpdatedAt string `gorm:"column:updated_at;not null"`
}
