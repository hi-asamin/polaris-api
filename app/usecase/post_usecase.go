package usecase

import (
	"mime/multipart"
	"polaris-api/infrastructure/repository"
	"polaris-api/interface/model"
)

type PostUseCase struct{}

func (u *PostUseCase) NewPost(userID, placeID, body string, files []*multipart.FileHeader) error {
	s3Repo := &repository.S3Repository{}
	placeRepo := &repository.PlaceRepository{}
	postRepo := &repository.PostRepository{}

	// 投稿対象の場所が存在するかチェック
	place, err := placeRepo.FindByID(placeID)
	if err != nil {
		return err
	}

	var fileInfos []model.FileInfo
	// S3にファイルをアップロード
	for _, file := range files {
		fileNameWithoutExt, err := s3Repo.UploadImage(file, placeID, userID)
		if err != nil {
			return err
		}

		// ファイルタイプを判定
		contentType := file.Header.Get("Content-Type")
		fileType := "image"
		if contentType == "video/mp4" || contentType == "video/quicktime" {
			fileType = "video"
		}

		fileInfos = append(fileInfos, model.FileInfo{
			FileName: fileNameWithoutExt,
			FileType: fileType,
		})
	}

	// TODO: デフォルトで公開設定（今後、下書き保存機能を実装したい）
	// 投稿の本文とメディア情報をデータベースに格納
	err = postRepo.CreatePost(userID, placeID, place.Name, body, true, fileInfos)
	if err != nil {
		return err
	}

	return nil
}

func (u *PostUseCase) DeletePost(postID string) error {
	postRepo := &repository.PostRepository{}

	// 投稿を削除
	err := postRepo.DeletePost(postID)
	if err != nil {
		return err
	}

	return nil
}
