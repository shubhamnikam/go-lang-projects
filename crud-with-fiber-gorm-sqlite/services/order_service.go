package services

import (
	"errors"

	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/models"
)

func GetAllOrders() ([]models.Order, error) {
	orders := []models.Order{}
	DbContext.Db.Find(&orders)
	if len(orders) == 0 {
		return orders, errors.New("order not found")
	}
	return orders, nil
}

func GetOrder(id int) (models.Order, error) {
	order := models.Order{}
	DbContext.Db.Find(&order, "id = ?", id)
	if order.Id == 0 {
		return order, errors.New("order not found with id=" + string(id))
	}

	return order, nil
}

func CreateOrder(order *models.Order) {
	DbContext.Db.Create(&order)
}

func DeleteOrder(id int) (models.Order, error) {
	//find
	order := models.Order{}
	DbContext.Db.Find(&order, "id=?", id)
	if order.Id == 0 {
		return order, errors.New("order not found with id=" + string(id))
	}
	//delete
	if err := DbContext.Db.Delete(&order).Error; err != nil {
		return order, errors.New("Error in deletion:" + err.Error())
	}

	return order, nil
}
