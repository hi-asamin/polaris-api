package router

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init はローカル環境用のルーターを初期化して起動します
func Init() {
	log.Print("Init router for local environment")

	r := CreateRouter()

	// サーバーを起動
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func CreateRouter() *gin.Engine {
	log.Print("Init router")

	r := gin.Default()

	// CORSミドルウェアの設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // フロントエンドのオリジンを許可
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// "api/v1" グループを作成
	g := r.Group("/api/v1")

	// Router
	PlaceRouter(g)
	CategoryRouter(g)

	return r
}
