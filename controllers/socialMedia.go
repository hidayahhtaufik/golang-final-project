package controllers

import (
	"myGram/database"
	"myGram/helpers"
	"myGram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm/clause"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("UserData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	socialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	socialMedia.UserID = userID

	err := db.Debug().Create(&socialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaURL,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
		"message":          "Success Created Social Media",
	})
}
func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("UserData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialMedia := []models.SocialMedia{}
	resData := []map[string]interface{}{}
	_ = resData
	err := db.Preload("User").Where("user_id=?", userID).Find(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for i := range socialMedia {
		nestedData := map[string]interface{}{
			"id":                socialMedia[i].User.ID,
			"username":          socialMedia[i].User.Username,
			"profile_image_url": "place holder string, di spesifikasi tabel tidak ada satupun kolom profile_image_url",
		}
		data := map[string]interface{}{
			"id":               socialMedia[i].ID,
			"name":             socialMedia[i].Name,
			"social_media_url": socialMedia[i].SocialMediaURL,
			"user_id":          socialMedia[i].UserID,
			"created_at":       socialMedia[i].CreatedAt,
			"updated_at":       socialMedia[i].UpdatedAt,
			"User":             nestedData,
		}

		resData = append(resData, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusOK,
		"social_medias": resData,
		"message":       "Success get data Social Media",
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	socialMedia := models.SocialMedia{}

	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	err := db.Model(&socialMedia).Where("id=?", socialMediaID).Clauses(clause.Returning{}).Updates(models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaURL,
		"user_id":          socialMedia.UserID,
		"updated_at":       socialMedia.UpdatedAt,
		"message":          "Success Update Social Media",
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaID"))

	var socialMedia models.SocialMedia

	err := db.Model(&socialMedia).Delete(&socialMedia, socialMediaID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "your social media has been sucessfully deleted",
	})
}
