package repository

import (
	"github.com/CHTTCH/little_project/backend/entity/order"
	"gorm.io/gorm"
)

type OrderRepository struct {
	dB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{dB: db}
}

func (repo *OrderRepository) Save(o *order.Order) error {
	if err := repo.dB.AutoMigrate(new(order.Order)); err != nil {
		return err
	}

	if err := repo.dB.Save(o).Error; err != nil {
		return err
	}

	return nil
}

func (repo *OrderRepository) FindAll() ([]order.Order, error) {
	orders := new([]order.Order)

	if err := repo.dB.Find(orders).Error; err != nil {
		return nil, err
	}

	return *orders, nil
}

func (repo *OrderRepository) FindById(id int) (*order.Order, error) {
	order := new(order.Order)

	if err := repo.dB.First(order, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return order, nil
}
