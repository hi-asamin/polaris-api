package usecase

import (
	"mime/multipart"
	"polaris-api/infrastructure/repository"
)

type PostUseCase struct{}

func (u *PostUseCase) NewPost(userID, placeID, body string, files []*multipart.FileHeader) error {
	s3Repo := &repository.S3Repository{}
	postRepo := &repository.PostRepository{}

	// S3にファイルをアップロード
	var fileNames []string
	for _, file := range files {
		fileNameWithoutExt, err := s3Repo.UploadImage(file, placeID, userID)
		if err != nil {
			return err
		}
		fileNames = append(fileNames, fileNameWithoutExt)
	}

	// TODO: デフォルトで公開設定（今後、下書き保存機能を実装したい）
	// 投稿の本文とメディア情報をデータベースに格納
	err := postRepo.CreatePost(userID, placeID, body, true, fileNames)
	if err != nil {
		return err
	}

	return nil
}
