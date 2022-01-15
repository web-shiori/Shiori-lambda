package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	tmpFileKey                  = "/tmp/tmp.png"
	tmpFile2Key                 = "/tmp/tmp2.png"
	targetS3Bucket              = "web-snapshot-inverted-pdf-screenshot"
	targetS3ObjectDirectoryName = "inverted_pdf/"
	targetS3ObjectFileName      = "/screenshot.png"
)

func downloadScreenshot(sess *session.Session, bucket string, key string) (*os.File, error) {
	fmt.Println("------スクリーンショットをS3からダウンロード------")
	f, err := os.Create(tmpFileKey)
	if err != nil {
		return nil, fmt.Errorf("%s: create tmp-file faild", err)
	}

	downloader := s3manager.NewDownloader(sess)
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if n == 0 || err != nil {
		return nil, fmt.Errorf("%s: open tmp-file2 faild", err)
	}

	return f, nil
}

func invert(file *os.File) (*image.RGBA, error) {
	fmt.Println("------画像の色を反転------")
	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("%s: decode faild", err)
	}
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	rgbaScale := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{w, h},
	})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imgColor := img.At(x, y)
			rr, gg, bb, aa := imgColor.RGBA()
			var max = int64(255)
			nr := max - int64(int8(rr))
			ng := max - int64(int8(gg))
			nb := max - int64(int8(bb))
			invertColor := color.RGBA{
				R: uint8(nr),
				G: uint8(ng),
				B: uint8(nb),
				A: uint8(aa),
			}
			rgbaScale.Set(x, y, invertColor)
		}
	}
	return rgbaScale, nil
}

func uploadInvertedScreenshot(sess *session.Session, uploadFile *os.File, contentID string) error {
	fmt.Println("------反転した画像をS3にアップロード------")
	uploader := s3manager.NewUploader(sess)
	key := targetS3ObjectDirectoryName + contentID + targetS3ObjectFileName
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(targetS3Bucket),
		Key:    aws.String(key),
		Body:   uploadFile,
	})
	if err != nil {
		return fmt.Errorf("%s: Upload faild", err)
	}
	return nil
}

// S3に保存されたオブジェクトからコンテンツIDを取得する。
// S3では"/pdf/コンテンツID/〇〇.png"というフォルダ構成になっている。
// ↑からコンテンツIDを取得する
func getContentID(key string) (string, error) {
	ss := strings.Split(key, "/")
	if len(ss) != 3 {
		return "", fmt.Errorf("%s: Invalid object", ss)
	}
	return ss[1], nil
}

func ImageColorInvert(event events.S3Event) error {
	sess := session.Must(session.NewSession())
	bucket := event.Records[0].S3.Bucket.Name
	key := event.Records[0].S3.Object.Key
	f, err := downloadScreenshot(sess, bucket, key)
	if err != nil {
		return err
	}

	rgba, err := invert(f)
	if err != nil {
		return err
	}

	nf, err := os.Create(tmpFile2Key)
	if err != nil {
		return fmt.Errorf("%s: create tmp-file2 faild", err)
	}

	err = png.Encode(nf, rgba)
	if err != nil {
		return fmt.Errorf("%s: open tmp-file2 faild", err)
	}

	uf, err := os.Open(tmpFile2Key)
	if err != nil {
		return fmt.Errorf("%s: open tmp-file2 faild", err)
	}

	id, err := getContentID(key)
	if err != nil {
		return fmt.Errorf("%s: get contentID faild", err)
	}

	err = uploadInvertedScreenshot(sess, uf, id)
	if err != nil {
		return err
	}

	return nil
}
