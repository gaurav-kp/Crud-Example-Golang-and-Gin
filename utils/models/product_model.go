package models

//define the model with the fields, type, handle unicode chars, size, etc

type ProductModel struct {
	ID          int    `gorm:"primaryKey"`
	Code        string `gorm:"size:20;column:product_code"`
	Price       uint   `gorm:"column:price"`
	Brand       string `gorm:"column:brand;type:varchar(100) CHARACTER SET utf8mb3"`
	Description string `gorm:"column:description;type:varchar(100) CHARACTER SET utf8mb3"`
}

// optional: use this to define the table name
func (ProductModel) TableName() string {
	return "product_tbl"
}
