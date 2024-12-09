package repository

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	"polaris-api/domain"
	"polaris-api/infrastructure"
)

type S3Repository struct{}

// UploadImageはS3にファイルをアップロードします
func (r *S3Repository) UploadImage(file *multipart.FileHeader, placeID, userID string) (string, error) {
	s3client := infrastructure.GetS3Client()
	bucketName := infrastructure.GetS3BucketName()

	src, err := file.Open()
	if err != nil {
		return "", domain.Wrap(err, 500, "S3へアップロードするファイルの参照時にエラーが発生")
	}
	defer src.Close()

	// ファイル名をユニークにする
	filename := fmt.Sprintf("%s/%s/%d%s", placeID, userID, time.Now().UnixNano(), filepath.Ext(file.Filename))
	contentType := file.Header.Get("Content-Type")

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(src); err != nil {
		return "", domain.Wrap(err, 500, "S3へアップロードするファイル読み取りエラー発生")
	}

	// S3にアップロード
	_, err = s3client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &filename,
		Body:        bytes.NewReader(buffer.Bytes()),
		ContentType: &contentType,
	})
	if err != nil {
		return "", domain.Wrap(err, 500, "S3へファイルアップロード時にエラー発生")
	}

	// 拡張子を取り除いたファイル名を返却（Lambdaで.webpに変換されるため）
	fileNameWithoutExt := filename[:len(filename)-len(filepath.Ext(filename))]
	return fileNameWithoutExt, nil
}
