package repository

import (
	customizedError "github.com/CHTTCH/little_project/backend/adapter/error"
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

	if err := repo.dB.Order("id asc").Save(o).Error; err != nil {
		return err
	}

	return nil
}

func (repo *OrderRepository) FindAll() ([]order.Order, error) {
	orders := new([]order.Order)

	if exist := repo.dB.Migrator().HasTable(orders); !exist {
		return nil, &customizedError.CustomizeError{Message: "table not exist."}
	} else if err := repo.dB.Find(orders).Error; err != nil {
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
