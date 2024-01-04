package app

import (
	"errors"
	"net/http"
	"tonx/pkg/data"
	"tonx/pkg/db"
	"tonx/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

// router.go

func CreateTestNetItem(c *gin.Context) {
	var newMetadata data.NftItemData
	if err := c.ShouldBindJSON(&newMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model := &models.NftTestnetItem{
		Name:        newMetadata.Name,
		Description: newMetadata.Description,
		Image:       newMetadata.Image,
		ExternalUrl: newMetadata.ExternalUrl,
		Marketplace: newMetadata.Marketplace,
	}

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

	c.JSON(http.StatusOK, gin.H{"message": " item created"})
}

func UpdateTestNetItem(c *gin.Context) {
	var updatedMetadata data.NftItemData
	if err := c.ShouldBindJSON(&updatedMetadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentItem = updatedMetadata // 更新 Metadata
	c.JSON(http.StatusOK, gin.H{"message": "item updated", "metadata": currentItem})
}

func GetTestNetItem(c *gin.Context) {
	var model models.NftTestnetItem
	// 假设我们使用 ID 字段来识别要获取的记录
	result := db.DB.First(&model, "name = ?", c.Param("name"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	metadata := &data.NftItemData{
		Name:        model.Name,
		Description: model.Description,
		Image:       model.Image,
		ExternalUrl: model.ExternalUrl,
		Marketplace: model.Marketplace,
	}

	c.JSON(http.StatusOK, metadata)
}
