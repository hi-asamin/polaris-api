package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"polaris-api/domain"
	"polaris-api/interface/handler"
)

func CategoryRouter(g *gin.RouterGroup) {
	log.Print("Category Router")
	categoryHandler := &handler.CategoryHandler{}

	g.GET("/categories", func(c *gin.Context) {
		// ハンドラー呼び出し
		categories, err := categoryHandler.GetCategories()
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
		c.JSON(http.StatusOK, gin.H{"categories": categories})
	})
}
