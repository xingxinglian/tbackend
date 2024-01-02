// router.go

package routers

import (
	"tonx/pkg/app"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 上传图片的接口
	router.POST("/upload", app.UploadHandler)

	// 更新和获取 Collection Metadata 的接口
	router.POST("/collection/create", app.CreateCollectionMetadata)
	router.GET("/collection", app.GetCollectionMetadata)

	// 更新和获取 TokenTick Metadata 的接口
	router.POST("/tokentick/create", app.CreateItem)
	router.GET("/tokentick", app.GetItem)

	return router
}
