// router.go

package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 上传图片的接口
	router.POST("/upload", uploadHandler)

	// 更新和获取 Collection Metadata 的接口
	router.POST("/collection/update", updateCollectionMetadata)
	router.GET("/collection", getCollectionMetadata)

	// 更新和获取 TokenTick Metadata 的接口
	router.POST("/tokentick/update", updateTokenTickMetadata)
	router.GET("/tokentick", getTokenTickMetadata)

	return router
}

func uploadHandler(c *gin.Context) {
	// 实现文件上传逻辑
}

func updateCollectionMetadata(c *gin.Context) {
	// 实现更新 Collection Metadata 逻辑
}

func getCollectionMetadata(c *gin.Context) {
	// 实现获取 Collection Metadata 逻辑
}

func updateTokenTickMetadata(c *gin.Context) {
	// 实现更新 TokenTick Metadata 逻辑
}

func getTokenTickMetadata(c *gin.Context) {
	// 实现获取 TokenTick Metadata 逻辑
}
