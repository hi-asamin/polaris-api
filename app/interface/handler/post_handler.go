package handler

import (
	"mime/multipart"
	"os"

	"polaris-api/domain"
	"polaris-api/usecase"
)

type PostHandler struct{}

func (h *PostHandler) NewPost(userID, placeID, body string, files []*multipart.FileHeader) error {
	u := &usecase.PostUseCase{}

	// ログインしていないユーザーが投稿した場合は、管理者アカウントとして投稿する
	if userID == "" {
		userID = os.Getenv("ADMIN_USER_ID")
	}

	// 場所IDが指定されていない場合、リクエストを受け付けない
	if placeID == "" {
		return domain.New(400, "場所IDの指定が不正")
	}

	if len(files) == 0 {
		return domain.New(400, "投稿するために必要なファイルが選択されていません")
	}
	if len(files) > 10 {
		return domain.New(400, "投稿できるファイルの上限数を超過")
	}

	err := u.NewPost(userID, placeID, body, files)
	if err != nil {
		return err
	}

	return nil
}

func (h *PostHandler) DeletePost(postID string) error {
	u := &usecase.PostUseCase{}

	if postID == "" {
		return domain.New(400, "投稿IDの指定が不正")
	}

	err := u.DeletePost(postID)
	if err != nil {
		return err
	}

	return nil
}
