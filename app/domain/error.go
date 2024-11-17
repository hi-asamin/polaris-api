package domain

import "log"

// AppError は独自のエラー型で、ステータスコードとメッセージを管理します
type AppError struct {
	Code    int
	Message string
	Err     error
}

// Error メソッドを追加して error インターフェースを実装
func (e *AppError) Error() string {
	log.Printf("%s: %v", e.Message, e.Err)
	return e.Message
}

// Wrap は既存のエラーをラップして AppError を作成します
func Wrap(err error, code int, message string) *AppError {
	log.Printf("%s: %v", message, err)
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// New は新しい AppError を作成します
func New(code int, message string) *AppError {
	log.Printf("%s", message)
	return &AppError{
		Code:    code,
		Message: message,
	}
}
