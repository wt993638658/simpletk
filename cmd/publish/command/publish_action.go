package command

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strings"

	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/wt993638658/simpletk/dal/db"
	"github.com/wt993638658/simpletk/kitex_gen/publish"
	"github.com/wt993638658/simpletk/pkg/minio"
	"github.com/wt993638658/simpletk/pkg/ttviper"
)

type PublishActionService struct {
	ctx context.Context
}

// NewPublishActionService new PublishActionService
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

// PublishAction publish video.
func (s *PublishActionService) PublishAction(req *publish.DouyinPublishActionRequest, uid int, cfg *ttviper.Config) (err error) {
	MinioVideoBucketName := minio.MinioVideoBucketName
	videoData := []byte(req.Data)

	// // 获取后缀
	// filetype := http.DetectContentType(videoData)
	fmt.Printf("开始存视频")
	// byte[] -> reader
	reader := bytes.NewReader(videoData)
	u2, err := uuid.NewV4()
	if err != nil {
		return err
	}
	fileName := u2.String() + "." + "mp4"
	// 上传视频
	fmt.Printf("上传视频\n")
	err = minio.UploadVideoFile(MinioVideoBucketName, fileName, reader, int64(len(videoData)))
	if err != nil {
		return err
	}
	// 获取视频链接
	fmt.Printf("获取视频连接")
	url, err := minio.GetFileUrl(MinioVideoBucketName, fileName, 0)
	fmt.Println("Url:%v\n", url)
	playUrl := strings.Split(url.String(), "?")[0]
	fmt.Println("playUrl:\n" + playUrl)
	if err != nil {
		return err
	}
	//fmt.Printf("\n" + playUrl + "\n")
	u3, err := uuid.NewV4()
	if err != nil {
		return err
	}
	fmt.Printf("获取封面")
	// 获取封面
	coverPath := u3.String() + "." + "jpg"
	fmt.Printf("开始从视频流截取一帧\n")

	//playUrl = "https://www.w3schools.com/html/movie.mp4"
	coverData, err := readFrameAsJpeg(playUrl)
	//coverData, err := readFrameAsJpeg(url)
	if err != nil {
		return err
	}
	fmt.Printf("上传封面")
	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err = minio.UploadJpgFile(MinioVideoBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		return err
	}
	fmt.Printf("获取封面链接")
	// 获取封面链接
	coverUrl, err := minio.GetFileUrl(MinioVideoBucketName, coverPath, 0)
	if err != nil {
		return err
	}

	CoverUrl := strings.Split(coverUrl.String(), "?")[0]
	////CoverUrl = "http://192.168.1.197" + strings.Split(CoverUrl, "localhost")[1]
	////playUrl = "http://192.168.1.197" + strings.Split(playUrl, "localhost")[1]
	//CoverUrl = "http://192.168.43.235" + strings.Split(CoverUrl, "localhost")[1]
	//playUrl = "http://192.168.43.235" + strings.Split(playUrl, "localhost")[1]
	////CoverUrl = "http://192.168.31.94" + strings.Split(CoverUrl, "localhost")[1]
	////playUrl = "http://192.168.31.94" + strings.Split(playUrl, "localhost")[1]
	videoModel := &db.Video{
		AuthorID:      uid,
		PlayUrl:       playUrl,
		CoverUrl:      CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}
	return db.CreateVideo(s.ctx, videoModel)
}

// ReadFrameAsJpeg
// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
	//func readFrameAsJpeg(filePath url.URL) ([]byte, error) {
	//fileBytes, _ := os.ReadFile("./data/bear.jpg")
	reader := bytes.NewBuffer(nil)

	fmt.Printf("开始行动")
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()

	fmt.Printf("结束行动")
	if err != nil {
		fmt.Printf("行动出错")
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}
