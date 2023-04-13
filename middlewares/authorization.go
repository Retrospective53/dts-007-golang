package middlewares

import (
	"challenge-10/database"
	"challenge-10/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}

		db.Where("id = ?", userID).First(&User)
		// skips this middleware if admin
		if User.Role == "ROLE_ADMIN" {
			c.Next()
			return
		}

		
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "bad request",
				"message": "invalid parameter",
			})
			return
		}
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data not found",
				"message" : "Data doesn't exist",
			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorized",
				"message" : "You are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}


func UserAdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}

		err := db.Where("id = ?", userID).First(&User).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data not found",
				"message" : "Data doesn't exist",
			})
			return
		}

		if User.Role != "ROLE_ADMIN" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorized",
				"message" : "You are not allowed to execute this operation",
			})
			return
		}

		c.Next()
	}
}