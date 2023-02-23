package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	ctx := context.Background()
	endpoint := "127.0.0.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "mymusic"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	objectName := "test.mp4"
	//filePath := "/Users/taowei/Desktop/go/tiktok-main/pkg/minio/client/test.mp4"
	//contentType := "application/zip"
	//
	//// Upload the zip file with FPutObject
	//info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	//object, err := minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer object.Close()
	//
	//localFile, err := os.Create("/tmp/test1.mp4")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer localFile.Close()
	//
	//if _, err = io.Copy(localFile, object); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	url, err := GetFileUrl(minioClient, bucketName, objectName, 0)
	playUrl := strings.Split(url.String(), "?")[0]
	fmt.Println("playUrl:" + playUrl)

}
func GetFileUrl(minioClient *minio.Client, bucketName string, fileName string, expires time.Duration) (*url.URL, error) {

	//fmt.Println(bucketName)
	//fmt.Println(fileName)
	ctx := context.Background()
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	fmt.Println("开始了\n")
	presignedUrl, err := minioClient.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		return nil, err
	}
	//fmt.Println("presignedUrl:%v", presignedUrl)
	fmt.Println("结束了\n")
	// TODO: url可能要做截取
	return presignedUrl, nil
}
