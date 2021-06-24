package site

type CreateSite struct {
	Website  string `json:"website" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateSite struct {
	Website  string `json:"website"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
