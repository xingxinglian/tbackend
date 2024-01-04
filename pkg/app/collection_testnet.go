package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"tonx/pkg/data"
	"tonx/pkg/db"
	"tonx/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateTestNetCollectionMetadata(c *gin.Context) {
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

	model := &models.CollectionTestnetMetadata{
		Name:        newMetadata.Name,
		Description: newMetadata.Description,
		Image:       newMetadata.Image,
		CoverImage:  newMetadata.CoverImage,
	}
	model.SocialLinks = string(links)

	// 使用 GORM 创建新的记录
	result := db.DB.Create(model)
	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == "23505" { // PostgreSQL 的唯一性违反错误代码
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Duplicated Name"})
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Collection Metadata created"})
}

func UpdateTestNetCollectionMetadata(c *gin.Context) {
	var updatedMetadata data.CollectionMetadata
	if err := c.ShouldBindJSON(&updatedMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model := &models.CollectionTestnetMetadata{
		Name:        updatedMetadata.Name,
		Description: updatedMetadata.Description,
		Image:       updatedMetadata.Image,
		CoverImage:  updatedMetadata.CoverImage,
		// SocialLinks: updatedMetadata.SocialLinks,
	}
	// model.SocialLinks = append(model.SocialLinks, updatedMetadata.SocialLinks...)

	// 假设我们使用 ID 字段来识别要更新的记录
	result := db.DB.Model(&models.CollectionTestnetMetadata{}).Where("name = ?", model.Name).Updates(model)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Collection Metadata updated", "metadata": updatedMetadata})
}

func GetTestNetCollectionMetadata(c *gin.Context) {
	var model models.CollectionTestnetMetadata
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
