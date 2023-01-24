package services

import "example.com/gin-api/models"

type Userservices interface {
	Createuser(*models.User) error
	Getuser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Updateuser(*models.User) error
	Deleteuser(*string) error
}
