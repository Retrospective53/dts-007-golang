package controllers

import (
	"challenge-10/database"
	"challenge-10/helpers"
	"challenge-10/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)

}
func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}
	
	productID, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productID)

	// var existingProduct models.Product
	// err := db.First(&existingProduct, productID).Error
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "product not found",
	// 	})
	// 	return
	// }


	err := db.Model(&Product).Where("id = ?", productID).Updates(models.Product{
		Title: Product.Title,
		Description: Product.Description,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func GetProductbyId(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}
	productID, _ := strconv.Atoi(c.Param("productId"))

	err := db.First(&Product, "id = ?", productID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, Product)
}

func GetProducts(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}
	products := []models.Product{}

	err := db.Model(&Product).Find(&products).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))

	err := db.Unscoped().Model(&Product).Where("id = ?", productId).Delete(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusAccepted)
}