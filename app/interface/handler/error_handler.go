package handler

import (
	"log"
	"net/http"

	"polaris-api/domain"

	"github.com/gin-gonic/gin"
)

// HandleError は共通のエラーレスポンス処理を行います
func HandleError(c *gin.Context, appErr *domain.AppError) {
	status := http.StatusInternalServerError

	// ステータスコードを AppError から取得
	switch appErr.Code {
	case 400:
		status = http.StatusBadRequest
	case 401:
		status = http.StatusUnauthorized
	case 403:
		status = http.StatusForbidden
	case 404:
		status = http.StatusNotFound
	case 500:
		status = http.StatusInternalServerError
	case 504:
		status = http.StatusGatewayTimeout
	default:
		log.Printf("HTTP %d: %s\n", http.StatusInternalServerError, "unknown error occurred")
		status = http.StatusInternalServerError
	}
	log.Printf("HTTP %d: %s\n", status, appErr.Message)

	// エラーログを出力
	c.JSON(status, gin.H{
		"error": appErr.Message,
	})
}
