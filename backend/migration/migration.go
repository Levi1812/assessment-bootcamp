package migration

type User struct {
	UserId   int `json:"primaryKey"`
	Name     string
	Address  string
	Email    string
	Password string
}
