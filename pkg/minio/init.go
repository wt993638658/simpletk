package minio

import (
	"context"
	"github.com/a76yyyy/tiktok/pkg/ttviper"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/prometheus/common/log"
)

var (
	minioClient          *minio.Client
	Config               = ttviper.ConfigInit("TIKTOK_MINIO", "minioConfig")
	MinioEndpoint        = Config.Viper.GetString("minio.Endpoint")
	MinioAccessKeyId     = Config.Viper.GetString("minio.AccessKeyId")
	MinioSecretAccessKey = Config.Viper.GetString("minio.SecretAccessKey")
	MinioUseSSL          = Config.Viper.GetBool("minio.UseSSL")
	MinioVideoBucketName = Config.Viper.GetString("minio.VideoBucketName")
)

// Minio 对象存储初始化
func init() {
	client, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKeyId, MinioSecretAccessKey, ""),
		Secure: MinioUseSSL,
	})
	if err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
	// fmt.Println(client)
	klog.Debug("minio client init successfully")
	minioClient = client
	if err := CreateBucket(MinioVideoBucketName); err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
	policy := `{"Version": "2012-10-17",
                "Statement": 
                    [{
                        "Action":["s3:GetObject"],
                        "Effect": "Allow",
                        "Principal": {"AWS": ["*"]},
                        "Resource": ["arn:aws:s3:::tiktok-video/*"],
                        "Sid": ""
                    }]
                }`
	ctx := context.Background()
	//action处设置读权限，resource处bucketname为桶的名称
	err = minioClient.SetBucketPolicy(ctx, MinioVideoBucketName, policy)
	if err != nil {
		log.Errorf("SetBucketPolicy   err:%s", err.Error())
	}
}
