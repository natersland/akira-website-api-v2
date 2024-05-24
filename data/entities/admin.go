package entities

import "time"

type Admin struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	UserName  string    `gorm:"column:username" json:"username"`
	Email     string    `gorm:"column:email;unique" json:"email" `
	Password  string    `gorm:"column:password" json:"password"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
