package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"polaris-api/domain"
	"polaris-api/interface/handler"
)

func PlaceRouter(g *gin.RouterGroup) {
	log.Print("Place Router")
	placeHandler := &handler.PlaceHandler{}

	g.GET("/places", func(c *gin.Context) {
		lonStr := c.Query("lon")
		latStr := c.Query("lat")
		cursorDistanceStr := c.Query("cursorDistance") // カーソル距離
		cursorPID := c.Query("cursorPID")              // カーソルのPlace ID
		cursorMID := c.Query("cursorMID")              // カーソルのMedia ID
		limitStr := c.Query("limit")                   // リミット件数

		// ハンドラー呼び出し
		response, err := placeHandler.GetPlaces(lonStr, latStr, cursorDistanceStr, cursorPID, cursorMID, limitStr)
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

	g.GET("/search", func(c *gin.Context) {
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
	// g.GET("/places/:id", func(c gin.Context) error { return placeController.Show(c) })
	// g.POST("/places", func(c gin.Context) error { return placeController.Create(c) })
	// g.PUT("/places/:id", func(c gin.Context) error { return placeController.Save(c) })
	// g.DELETE("/places/:id", func(c gin.Context) error { return placeController.Delete(c) })
}
