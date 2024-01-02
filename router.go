// router.go

package main

import (
	"tonx/pkg/app"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 上传图片的接口
	router.POST("/upload", app.UploadHandler)

	// 更新和获取 Collection Metadata 的接口
	router.POST("/collection/create", app.CreateCollectionMetadata)
	router.GET("/collection/get", app.GetCollectionMetadata)

	// 更新和获取 TokenTick Metadata 的接口
	router.POST("/tokentick/create", app.CreateItem)
	router.GET("/tokentick/get", app.GetItem)

	return router
}
