package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"polaris-api/domain"
	"polaris-api/interface/dto"
	"polaris-api/interface/handler"
)

func PlaceRouter(g *gin.RouterGroup) {
	log.Print("Place Router")
	placeHandler := &handler.PlaceHandler{}

	g.GET("/places", func(c *gin.Context) {
		cursorPID := c.Query("cursorPID") // カーソルのPlace ID
		cursorMID := c.Query("cursorMID") // カーソルのMedia ID
		limitStr := c.Query("limit")      // リミット件数

		// ハンドラー呼び出し
		response, err := placeHandler.GetPlaces(cursorPID, cursorMID, limitStr)
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

	g.POST("/places", func(c *gin.Context) {
		var req dto.CreatePlaceRequest

		// リクエストボディのバインドとGinによる初期バリデーション
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Print(err)
			// エラー処理を共通関数に委譲
			if appErr, ok := err.(*domain.AppError); ok {
				handler.HandleError(c, appErr)
			} else {
				handler.HandleError(c, domain.New(500, "Unknown error occurred"))
			}
			return
		}

		// ハンドラー呼び出し
		err := placeHandler.NewPlace(&req)
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

	g.GET("/places/:id", func(c *gin.Context) {
		// パスパラメータからIDを取得
		id := c.Param("id")

		// ハンドラー呼び出し
		place, err := placeHandler.GetPlacesByID(id)
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
		c.JSON(http.StatusOK, place)
	})

	g.GET("/places/:id/nearby", func(c *gin.Context) {
		// パスパラメータからIDを取得
		id := c.Param("id")
		lonStr := c.Query("lon")
		latStr := c.Query("lat")
		limitStr := c.Query("limit")

		// ハンドラー呼び出し
		response, err := placeHandler.GetPlacesNearBySpots(id, lonStr, latStr, limitStr)
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
}
