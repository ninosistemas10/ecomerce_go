package product

import (
	"github.com/ninosistemas10/ecommerce/model"

	"github.com/google/uuid"
)

type UseCase interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Product, error)
	GetAll() (model.Products, error)
	//GetPaginatedData(startIndex, limit int) (model.Products, error)
}

type Storage interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Product, error)
	GetAll() (model.Products, error)
	//GetPaginatedData(startIndex, limit int) (model.Products, error)
}
