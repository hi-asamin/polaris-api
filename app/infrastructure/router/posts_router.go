package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"polaris-api/domain"
	"polaris-api/interface/handler"
)

func PostsRouter(g *gin.RouterGroup) {
	log.Print("Posts Router")
	postHandler := &handler.PostHandler{}

	g.POST("/posts", func(c *gin.Context) {
		// ユーザーIDと場所IDを取得
		userID := c.PostForm("userId")
		placeID := c.PostForm("placeId")
		body := c.PostForm("body")

		// ファイルを受け取る
		form, err := c.MultipartForm()
		if err != nil {
			handler.HandleError(c, domain.New(400, "メディア情報が存在しない"))
			return
		}

		files := form.File["media"] // "media" フィールドにアップロードされたファイルを取得

		// ハンドラー呼び出し
		err = postHandler.NewPost(userID, placeID, body, files)
		if err != nil {
			// エラー処理を共通関数に委譲
			if appErr, ok := err.(*domain.AppError); ok {
				handler.HandleError(c, appErr)
			} else {
				handler.HandleError(c, domain.New(500, "Unknown error occurred"))
			}
			return
		}

		// 正常レスポンス
		c.JSON(http.StatusCreated, nil)
	})

}
