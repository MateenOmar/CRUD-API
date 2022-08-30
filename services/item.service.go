package services

import (
	"example/CRUD-APIs/models"
)

type ItemService interface {
	CreateItem(*models.Product) error
	GetItem(*string) (*models.Product, error)
	GetStock() ([]*models.Product, error)
	Purchase(*models.Product) error
	//Return(*models.Product) error
	DeleteItem(*string) error
}
