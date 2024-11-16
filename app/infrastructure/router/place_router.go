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

	g.GET("/search", func(c *gin.Context) {
		q := c.Query("q")
		log.Printf("%s", q)
		lonStr := c.Query("lon")
		latStr := c.Query("lat")
		places, err := placeHandler.GetPlaceByName(q, lonStr, latStr)
		if err != nil {
			// エラー処理を共通関数に委譲
			if appErr, ok := err.(*domain.AppError); ok {
				handler.HandleError(c, appErr)
			} else {
				handler.HandleError(c, domain.New(500, "Unknown error occurred"))
			}
			return
		}
		log.Printf("%s\n", places)

		c.JSON(http.StatusOK, gin.H{"places": places})
	})
	// g.GET("/places/:id", func(c gin.Context) error { return placeController.Show(c) })
	// g.POST("/places", func(c gin.Context) error { return placeController.Create(c) })
	// g.PUT("/places/:id", func(c gin.Context) error { return placeController.Save(c) })
	// g.DELETE("/places/:id", func(c gin.Context) error { return placeController.Delete(c) })
}