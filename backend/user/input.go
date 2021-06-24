package user

type UserRegisterinput struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogininput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
