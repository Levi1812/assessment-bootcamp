package site

import (
	"fmt"
	"time"
)

type Service interface {
	GetAllSiteByUser(UserId string) ([]Site, error)
	CreateSite(UserId int, input CreateSite) (Site, error)
	GetSiteByID(SiteId string) (Site, error)
	UpdateSite(SiteId string, input UpdateSite) (Site, error)
	DeleteSite(SiteId string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// Service

func (s *service) GetAllSiteByUser(UserId string) ([]Site, error) {
	site, err := s.repository.FindByUserID(UserId)

	if err != nil {
		return site, err
	}

	return site, nil
}

func (s *service) CreateSite(UserId int, input CreateSite) (Site, error) {
	var newSite = Site{
		UserId:    UserId,
		Website:   input.Website,
		Login:     input.Login,
		Password:  input.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	site, err := s.repository.Create(newSite)

	if err != nil {
		return site, err
	}

	return site, nil
}

func (s *service) GetSiteByID(SiteId string) (Site, error) {
	site, err := s.repository.FindByID(SiteId)

	if err != nil {
		return site, err
	}

	return site, nil
}

func (s *service) UpdateSite(SiteId string, input UpdateSite) (Site, error) {
	var dataUpdate = map[string]interface{}{}

	if input.Website != "" || len(input.Website) != 0 {
		dataUpdate["website"] = input.Website
	}

	if input.Login != "" || len(input.Login) != 0 {
		dataUpdate["login"] = input.Login
	}

	if input.Password != "" || len(input.Password) != 0 {
		dataUpdate["password"] = input.Password
	}

	dataUpdate["updated_at"] = time.Now()

	site, err := s.repository.Update(SiteId, dataUpdate)

	if err != nil {
		return site, err
	}

	return site, nil
}

func (s *service) DeleteSite(SiteId string) (string, error) {
	site, err := s.repository.Delete(SiteId)

	if err != nil || site == "Error" {
		return site, err
	}

	message := fmt.Sprintf("Password pada Id %s berhasil dihapus", SiteId)
	return message, nil
}
