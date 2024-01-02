package app

import (
	"encoding/json"
	"net/http"

	"tonx/pkg/data"
	"tonx/pkg/db"
	"tonx/pkg/models"

	"github.com/gin-gonic/gin"
)

func CreateCollectionMetadata(c *gin.Context) {
	var newMetadata data.CollectionMetadata
	if err := c.ShouldBindJSON(&newMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// json.Marshal
	links, err := json.Marshal(newMetadata.SocialLinks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	model := &models.CollectionMetadata{
		Name:        newMetadata.Name,
		Description: newMetadata.Description,
		Image:       newMetadata.Image,
		CoverImage:  newMetadata.CoverImage,
	}
	model.SocialLinks = string(links)

	// 使用 GORM 创建新的记录
	result := db.DB.Create(model)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Collection Metadata created"})
}

func UpdateCollectionMetadata(c *gin.Context) {
	var updatedMetadata data.CollectionMetadata
	if err := c.ShouldBindJSON(&updatedMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model := &models.CollectionMetadata{
		Name:        updatedMetadata.Name,
		Description: updatedMetadata.Description,
		Image:       updatedMetadata.Image,
		CoverImage:  updatedMetadata.CoverImage,
		// SocialLinks: updatedMetadata.SocialLinks,
	}
	// model.SocialLinks = append(model.SocialLinks, updatedMetadata.SocialLinks...)

	// 假设我们使用 ID 字段来识别要更新的记录
	result := db.DB.Model(&models.CollectionMetadata{}).Where("name = ?", model.Name).Updates(model)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Collection Metadata updated", "metadata": updatedMetadata})
}

func GetCollectionMetadata(c *gin.Context) {
	var model models.CollectionMetadata
	// 假设我们使用 ID 字段来识别要获取的记录
	result := db.DB.First(&model, "name = ?", c.Param("name"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	metadata := &data.CollectionMetadata{
		Name:        model.Name,
		Description: model.Description,
		Image:       model.Image,
		CoverImage:  model.CoverImage,
		// SocialLinks: model.SocialLinks,
	}
	metadata.SocialLinks = make([]string, 0)
	// var sss interface{}
	err := json.Unmarshal([]byte(model.SocialLinks), &metadata.SocialLinks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, metadata)
}
