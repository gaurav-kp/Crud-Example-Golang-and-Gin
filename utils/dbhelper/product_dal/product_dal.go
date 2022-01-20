package product_dal

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	constants "microservice/utils/dbhelper"
	m "microservice/utils/models"
)

func Save(product m.ProductModel) (m.ProductModel, error) {

	db, err := gorm.Open(mysql.Open(constants.DsnMySql), &gorm.Config{})

	if err != nil {
		return product, err
	}

	// Migrate the schema
	db.AutoMigrate(&m.ProductModel{})

	// pass pointer of data to save Product
	result := db.Create(&product)

	// confirm if the row is saved and return
	if result.RowsAffected > 0 {
		return product, nil
	}

	return product, result.Error
}

func Update(product m.ProductModel) (m.ProductModel, error) {
	db, err := gorm.Open(mysql.Open(constants.DsnMySql), &gorm.Config{})

	if err != nil {
		return product, err
	}

	// pass pointer of data to update Product
	updateRes := db.Where("id = ?", product.ID).UpdateColumns(&product)

	//check if rows affected is greater than 0
	if updateRes.RowsAffected > 0 {
		return product, nil
	}

	return product, errors.New("no rows affected")

}

func Delete(id int) int64 {
	db, err := gorm.Open(mysql.Open(constants.DsnMySql), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// check if the product exists
	var product m.ProductModel
	getRes := db.First(&product, "id = ?", id)

	if getRes.RowsAffected < 1 {
		println("error cannot delete record, no record found")
		return -1
	}

	// this will will delete product based on id, primary key
	deleteRes := db.Delete(&product)

	if deleteRes.Error != nil {
		return -1
	}

	return deleteRes.RowsAffected
}

func GetAll() ([]m.ProductModel, error) {

	db, err := gorm.Open(mysql.Open(constants.DsnMySql), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var products []m.ProductModel

	// pass pointer of data to get all products
	result := db.Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func GetOne(id int) (m.ProductModel, error) {

	var product m.ProductModel

	db, err := gorm.Open(mysql.Open(constants.DsnMySql), &gorm.Config{})

	if err != nil {
		return product, err
	}

	// pass pointer of data to get product based on id
	result := db.First(&product, "id = ?", id)

	// return of no record found
	if result.RowsAffected < 1 {
		println("no record found")
	}

	return product, nil
}
