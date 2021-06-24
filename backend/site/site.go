package site

import "time"

type Site struct {
	SiteId    int       `json:"site_id" gorm:"primaryKey"`
	Website   string    `json:"website"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    int       `json:"user_id"`
}
