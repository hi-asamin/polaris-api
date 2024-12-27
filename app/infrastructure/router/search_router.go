package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"polaris-api/domain"
	"polaris-api/interface/handler"
)

func SearchRouter(g *gin.RouterGroup) {
	log.Print("Search Router")
	placeHandler := &handler.PlaceHandler{}

	g.GET("/search", func(c *gin.Context) {
		keywords := c.Query("keywords")   // 検索ワード
		cursorMID := c.Query("cursorMID") // カーソルのMedia ID
		limitStr := c.Query("limit")      // リミット件数

		// ハンドラー呼び出し
		response, err := placeHandler.GetPlacesBaseQuery(keywords, cursorMID, limitStr)
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
		c.JSON(http.StatusOK, response)
	})

	g.GET("/search/suggests", func(c *gin.Context) {
		q := c.Query("q")
		lonStr := c.Query("lon")
		latStr := c.Query("lat")
		places, err := placeHandler.GetPlacesByName(q, lonStr, latStr)
		if err != nil {
			// エラー処理を共通関数に委譲
			if appErr, ok := err.(*domain.AppError); ok {
				handler.HandleError(c, appErr)
			} else {
				handler.HandleError(c, domain.New(500, "Unknown error occurred"))
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"places": places})
	})

}
