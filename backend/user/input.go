package migration

type UserRegisterinput struct {
	Name     string `json:"name" binding:"requiered"`
	Address  string `json:"address" binding:"requiered"`
	Email    string `json:"email" binding:"requiered, email"`
	Password string `json:"password" binding:"requiered"`
}

type UserLogininput struct {
	Email    string `json:"email" binding:"requiered, email"`
	Password string `json:"password" binding:"requiered"`
}
