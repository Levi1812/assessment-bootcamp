package migration

type UserFormatter struct {
	UserId  int    `json:"userid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func UserFormat(migration User) UserFormatter {
	return UserFormatter{
		UserId:  migration.UserId,
		Name:    migration.Name,
		Email:   migration.Email,
		Address: migration.Address,
	}
}

func UserFormats(users []User) []UserFormatter {
	var userFormats []UserFormatter

	for _, user := range users {
		userFormats = append(userFormats, UserFormat(user))
	}

	return userFormats
}
