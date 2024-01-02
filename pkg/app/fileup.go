package app

import (
	"fmt"
	"net/http"
	"tonx/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

var (
	S3_REGION     = config.Config.Aws.S3Region
	S3_BUCKET     = config.Config.Aws.S3Bucket
	S3_KEY_ID     = config.Config.Aws.S3KeyId
	S3_ACCESS_KEY = config.Config.Aws.S3AccessKey
)

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 读取文件
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	// 创建一个 AWS session
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(S3_REGION),
		Credentials: credentials.NewStaticCredentials(S3_KEY_ID, S3_ACCESS_KEY, ""),
	})

	// 创建一个 S3 服务客户端
	s3Client := s3.New(sess)

	// 上传到 S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(S3_BUCKET),
		Key:                aws.String(file.Filename),
		ACL:                aws.String("public-read"),
		Body:               src,
		ContentDisposition: aws.String("attachment"),
		ContentType:        aws.String(http.DetectContentType([]byte(file.Filename))),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": aerr.Error()})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", S3_BUCKET, file.Filename)
	c.JSON(http.StatusOK, gin.H{"url": url})
}
