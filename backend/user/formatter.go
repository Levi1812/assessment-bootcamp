package user

type UserFormatter struct {
	UserId  int    `json:"userid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func UserFormat(user User) UserFormatter {
	return UserFormatter{
		UserId:  user.UserId,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}
}

func UserFormats(users []User) []UserFormatter {
	var userFormats []UserFormatter

	for _, user := range users {
		userFormats = append(userFormats, UserFormat(user))
	}

	return userFormats
}
