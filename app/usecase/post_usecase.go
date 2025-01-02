package usecase

import (
	"mime/multipart"
	"polaris-api/infrastructure/repository"
	"polaris-api/interface/types"
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

	var fileInfos []types.FileInfo
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

		fileInfos = append(fileInfos, types.FileInfo{
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
	s3Repo := &repository.S3Repository{}
	mediaRepo := &repository.MediaRepository{}
	postRepo := &repository.PostRepository{}

	// 投稿に紐づくメディアを削除
	medias, err := mediaRepo.FindByPostID(postID)
	if err != nil {
		return err
	}

	// メディアをS3から削除
	for _, media := range medias {
		key := media.MediaURL
		if media.MediaType == "image" {
			key = key + ".webp"
		}
		//else if media.MediaType == "video" {
		//	key = key + "_compressed.mp4"
		//}
		err = s3Repo.DeleteMedia(key)
		if err != nil {
			return err
		}
	}

	// 投稿を削除
	err = postRepo.DeletePost(postID)
	if err != nil {
		return err
	}

	return nil
}
