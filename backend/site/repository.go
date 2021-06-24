package site

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Site, error)
	FindByID(SiteId string) (Site, error)
	FindByUserID(UserId string) ([]Site, error)
	Create(site Site) (Site, error)
	Update(SiteId string, dataUpdate map[string]interface{}) (Site, error)
	Delete(SiteId string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

// Repository

func (r *repository) GetAll() ([]Site, error) {
	var site []Site

	if err := r.db.Find(&site).Error; err != nil {
		return site, err
	}

	return site, nil
}

func (r *repository) FindByID(SiteId string) (Site, error) {
	var site Site

	if err := r.db.Where("site_id = ?", SiteId).Find(&site).Error; err != nil {
		return site, err
	}

	return site, nil
}

func (r *repository) FindByUserID(UserId string) ([]Site, error) {
	var site []Site

	if err := r.db.Where("user_id = ?", UserId).Find(&site).Error; err != nil {
		return site, err
	}

	return site, nil
}

func (r *repository) Create(site Site) (Site, error) {
	if err := r.db.Create(&site).Error; err != nil {
		return site, err
	}

	return site, nil
}

func (r *repository) Update(SiteId string, dataUpdate map[string]interface{}) (Site, error) {
	var site Site

	if err := r.db.Model(&site).Where("site_id = ?", SiteId).Updates(dataUpdate).Error; err != nil {
		return site, err
	}

	if err := r.db.Where("site_id = ?", SiteId).Find(&site).Error; err != nil {
		return site, err
	}

	return site, nil
}

func (r *repository) Delete(SiteId string) (string, error) {
	var site Site

	if err := r.db.Where("site_id = ?", SiteId).Delete(&site).Error; err != nil {
		return "error delete", err
	}

	return "success", nil
}
