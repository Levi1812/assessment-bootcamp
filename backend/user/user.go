package user

import (
	"assessment/site"
	"time"
)

type User struct {
	UserId    int         `json:"user_id" gorm:"primaryKey"`
	Name      string      `json:"name"`
	Address   string      `json:"address"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Site      []site.Site `json:"site" gorm:foreignKey`
}
