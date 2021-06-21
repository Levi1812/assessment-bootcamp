package users

import "time"

type User struct {
	UserId    int `json:"primaryKey"`
	Name      string
	Address   string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
