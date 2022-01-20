package product_controller

import (
	product_dal "microservice/utils/DbHelper/product_dal"
	models "microservice/utils/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// register routes
func SetApiRoutes(r *gin.Engine) {
	r.POST("/create_product", CreateProduct)
	r.PUT("/update_product", UpdateProduct)
	r.DELETE("/delete_product", DeleteProduct)
	r.GET("/get_products", GetProducts)
	r.GET("/get_product", GetProduct)
}

func CreateProduct(c *gin.Context) {
	var json models.ProductModel

	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	data, err := product_dal.Save(json)

	println("CreateProduct  error -- ", err)

	if err == nil {
		c.AsciiJSON(http.StatusCreated, data)
	} else {
		c.AsciiJSON(http.StatusBadRequest, "error, unable to save product")
	}
}

func UpdateProduct(c *gin.Context) {
	var json models.ProductModel

	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	data, err := product_dal.Update(json)

	if err == nil {
		c.AsciiJSON(http.StatusOK, data)
	} else {
		c.AsciiJSON(http.StatusBadRequest, "error, unable to update product")
	}
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	rowsAffected := product_dal.Delete(id)

	if rowsAffected > 0 {
		c.AsciiJSON(http.StatusOK, rowsAffected)
	} else {
		// return no content
		c.AsciiJSON(http.StatusBadRequest, rowsAffected)
	}
}

func GetProducts(c *gin.Context) {

	data, err := product_dal.GetAll()

	if err == nil {
		c.AsciiJSON(http.StatusOK, data)
	} else {
		c.AsciiJSON(http.StatusNoContent, nil)
	}
}

func GetProduct(c *gin.Context) {

	id, _ := strconv.Atoi(c.Query("id"))

	data, err := product_dal.GetOne(id)

	if err == nil && data.ID > 0 {
		c.AsciiJSON(http.StatusOK, data)
	} else {
		// return no content
		c.AsciiJSON(http.StatusNoContent, nil)
	}
}
