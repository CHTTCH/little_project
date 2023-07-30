package repository

import (
	"github.com/CHTTCH/little_project/backend/entity/order"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (repo *OrderRepository) Save(o *order.Order) error {
	if err := repo.DB.AutoMigrate(new(order.Order)); err != nil {
		return err
	}

	if err := repo.DB.Save(o).Error; err != nil {
		return err
	}

	return nil
}

func (repo *OrderRepository) FindAll() ([]order.Order, error) {
	orders := new([]order.Order)

	if err := repo.DB.Find(orders).Error; err != nil {
		return nil, err
	}

	return *orders, nil
}

// TODO not done
func (repo *OrderRepository) FindById(id int) (*order.Order, error) {
	patients := new([]order.Order)

	if err := repo.DB.Find(patients).Error; err != nil {
		return nil, err
	}

	return new(order.Order), nil
}
