package repository

import (
	"github.com/CHTTCH/little_project/backend/entity/order"
	"github.com/CHTTCH/little_project/backend/entity/patient"
)

type Repository[T patient.Patient | order.Order, U int | string] interface {
	Save(*T) error
	FindAll() ([]T, error)
	FindById(U) (*T, error)
}
