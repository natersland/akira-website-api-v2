package entities

import "time"

type Admin struct {
	ID        string    `gorm:"column:id;primaryKey"`
	UserName  string    `gorm:"type:varchar(256);not null;column:username"`
	Email     string    `gorm:"type:varchar(256);not null;column:email"`
	Password  string    `gorm:"type:varchar(256);not null;column:password"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;column:updated_at"`
}
