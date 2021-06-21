package user

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]User, error)
	FindByID(UserId string) (User, error)
	FindByEmail(Email string) (User, error)
	Create(user User) (User, error)
	Update(UserId string, userUpdate map[string]interface{}) (User, error)
	Delete(UserId string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

// REPOSITORY

func (r *repository) GetAll() ([]User, error) {
	var users []User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindByID(UserId string) (User, error) {
	var user User

	if err := r.db.Where("id = ?", UserId).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(Email string) (User, error) {
	var user User

	if err := r.db.Where("email = ?", Email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Create(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(UserId string, userUpdate map[string]interface{}) (User, error) {
	var user User

	if err := r.db.Model(&user).Where("id = ?", UserId).Updates(userUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", UserId).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Delete(UserId string) (string, error) {
	var user User

	if err := r.db.Where("id = ?", UserId).Delete(&user).Error; err != nil {
		return "404", err
	}

	return "Success", nil
}
