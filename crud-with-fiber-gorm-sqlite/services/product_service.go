package services

import (
	"errors"

	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/models"
)

func GetAllProducts() ([]models.Product){
	products:= []models.Product{}
	DbContext.Db.Find(&products)
	return products
}

func GetProduct(id int) (models.Product, error) {
	product := models.Product{}
	DbContext.Db.Find(&product, "id = ?", id)
	if product.Id == 0 {
		return product, errors.New("product not found with id=" + string(id)) 
	} 

	return product, nil
}

func CreateProduct(product *models.Product) {
	DbContext.Db.Create(&product)
}

func UpdateProduct(id int, inputProduct models.Product) (models.Product, error) {
	//find 
	product := models.Product{}
	DbContext.Db.Find(&product, "id = ?", id)
	if product.Id == 0 {
		return product, errors.New("product not found with id=" + string(id)) 
	} 
	//update
	product.Name = inputProduct.Name
	product.SerialNo = inputProduct.SerialNo
	DbContext.Db.Save(&product)
	if err := DbContext.Db.Save(&product).Error; err != nil {
		return product, errors.New("Error in update:"+err.Error())
	}

	return product, nil
}

func DeleteProduct (id int) (models.Product, error) {
	//find
	product := models.Product{}
	DbContext.Db.Find(&product, "id=?", id)
	if product.Id == 0 {
		return product, errors.New("product not found with id="+string(id))
		} 
	//delete
	if err := DbContext.Db.Delete(&product).Error; err != nil {
		return product, errors.New("Error in deletion:"+err.Error())
	}
	
	return product, nil
}