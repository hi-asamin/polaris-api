package router

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"polaris-api/domain"
	"polaris-api/interface/handler"
	"polaris-api/interface/types"
)

func PlaceRouter(g *gin.RouterGroup) {
	log.Print("Place Router")
	placeHandler := &handler.PlaceHandler{}

	g.GET("/places", func(c *gin.Context) {
		cursorMID := c.Query("cursorMID") // カーソルのMedia ID
		limitStr := c.Query("limit")      // リミット件数

		var categoryIds []int
		categoryIdsStr := c.Query("categoryIds")
		if categoryIdsStr != "" {
			// カンマ区切りの文字列を分割
			categoryIdsStrArr := strings.Split(categoryIdsStr, ",")
			categoryIds = make([]int, 0, len(categoryIdsStrArr))

			// 各文字列を数値に変換
			for _, idStr := range categoryIdsStrArr {
				if id, err := strconv.Atoi(idStr); err == nil {
					categoryIds = append(categoryIds, id)
				}
			}
		}

		// ハンドラー呼び出し
		response, err := placeHandler.GetPlaces(cursorMID, limitStr, categoryIds)
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
		var req types.CreatePlaceRequest

		// リクエストボディのバインドとGinによる初期バリデーション
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Print(err)
			// エラー処理を共通関数に委譲
			if appErr, ok := err.(*domain.AppError); ok {
				handler.HandleError(c, appErr)
			} else {
				handler.HandleError(c, domain.New(500, "バリデーションエラー"))
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

	g.PUT("/places/:id", func(c *gin.Context) {
		// パスパラメータからIDを取得
		id := c.Param("id")

		// JSONボディを受け取る
		var payload types.PlaceUpdatePayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			handler.HandleError(c, domain.New(400, "Invalid request body"))
			return
		}

		err := placeHandler.UpdatePlace(id, &payload)
		if err != nil {
			// エラー処理を共通関数に委譲
			if appErr, ok := err.(*domain.AppError); ok {
				handler.HandleError(c, appErr)
			} else {
				handler.HandleError(c, domain.New(500, "Unknown error occurred"))
			}
			return
		}

		c.JSON(http.StatusNoContent, nil)
	})

	g.GET("/places/:id/nearby", func(c *gin.Context) {
		// パスパラメータからIDを取得
		id := c.Param("id")
		lonStr := c.Query("lon")
		latStr := c.Query("lat")
		cursorMID := c.Query("cursorMID")
		limitStr := c.Query("limit")

		// ハンドラー呼び出し
		response, err := placeHandler.GetPlacesNearBySpots(id, lonStr, latStr, cursorMID, limitStr)
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
