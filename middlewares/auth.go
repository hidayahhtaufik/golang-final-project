package middlewares

import (
	"myGram/database"
	"myGram/helpers"
	"myGram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"error":   "Unauthorized",
				"message": "Not Authentication",
			})
			return
		}
		c.Set("UserData", verifyToken)
		c.Next()
	}
}

func Authorization(input string) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoID, err := strconv.Atoi(c.Param("photoID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "Bad Request",
				"message": "Invalid Params",
			})
			return
		}

		userData := c.MustGet("UserData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"error":   "Data Not Found",
				"message": "Data Doesn't Exist",
			})
			return
		}

		if Photo.User.ID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"error":   "Unauthorized",
				"message": "Not Authorized",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socMedID, err := strconv.Atoi(c.Param("socialMediaID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "Bad Request",
				"message": "Invalid Params",
			})
			return
		}

		userData := c.MustGet("UserData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		socMed := models.SocialMedia{}

		err = db.Select("user_id").First(&socMed, uint(socMedID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"error":   "Data Not Found",
				"message": "Data Doesn't Exist",
			})
			return
		}
		if socMed.UserID != userID {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"message": "Not Authorized",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentID, err := strconv.Atoi(c.Param("commentID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "Bad Request",
				"message": "Invalid Params",
			})
			return
		}

		userData := c.MustGet("UserData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := models.Comment{}

		err = db.Select("user_id").First(&comment, uint(commentID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"error":   "Data Not Found",
				"message": "Data Doesn't Exist",
			})
			return
		}

		if comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"error":   "Unauthorized",
				"message": "Not Authorized",
			})
			return
		}

		c.Next()
	}
}
