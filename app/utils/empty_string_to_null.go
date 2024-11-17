package utils

// 空文字列を NULL に変換するヘルパー関数
func EmptyStringToNull(value string) interface{} {
	if value == "" {
		return nil // NULL を返す
	}
	return value // そのまま値を返す
}
