package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input UserRegisterinput) (UserFormatter, error)
	LoginUser(input UserLogininput) (UserFormatter, error)
	GetAllUser() ([]UserFormatter, error)
	GetUserByID(UserId string) (UserFormatter, error)
	UpdateUser(UserId string, input UpdateUser) (UserFormatter, error)
	DeleteUser(UserId string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// Service

func (s *service) RegisterUser(input UserRegisterinput) (UserFormatter, error) {
	generatePass, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser = User{
		Name:      input.Name,
		Address:   input.Address,
		Email:     input.Email,
		Password:  string(generatePass),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.repository.Create(newUser)
	if err != nil {
		return UserFormatter{}, err
	}
	formatter := UserFormat(user)
	return formatter, nil
}

func (s *service) LoginUser(input UserLogininput) (UserFormatter, error) {

	checkUser, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return UserFormatter{}, err
	}
	if checkUser.UserId == 0 || len(checkUser.Name) <= 1 {
		return UserFormatter{}, errors.New("MASUKAN EMAIL / PASSWORD DENGAN BENAR")
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(input.Password))
	if err != nil {
		return UserFormatter{}, errors.New("MASUKAN EMAIL / PASSWORD DENGAN BENAR")
	}

	formatter := UserFormat(checkUser)

	return formatter, nil

}

func (s *service) GetAllUser() ([]UserFormatter, error) {
	users, err := s.repository.GetAll()

	if err != nil {
		return []UserFormatter{}, nil
	}

	formatters := UserFormats(users)

	return formatters, nil
}

func (s *service) GetUserByID(UserId string) (UserFormatter, error) {
	user, err := s.repository.FindByID(UserId)

	if err != nil {
		return UserFormatter{}, nil
	}

	formatter := UserFormat(user)

	return formatter, nil
}

func (s *service) UpdateUser(UserId string, input UpdateUser) (UserFormatter, error) {
	var userUpdate = map[string]interface{}{}

	if input.Name != "" || len(input.Name) != 0 {
		userUpdate["Name"] = input.Name
	}

	if input.Address != "" || len(input.Address) != 0 {
		userUpdate["Address"] = input.Address
	}

	userUpdate["updated_at"] = time.Now()

	user, err := s.repository.Update(UserId, userUpdate)

	if err != nil {
		return user, err
	}

	return user, nil
}
