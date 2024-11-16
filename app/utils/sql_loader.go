package utils

import (
    "fmt"
    "io/ioutil"
)

// LoadSQLFile は指定したパスのSQLファイルを読み込み、その内容を文字列として返します。
func LoadSQLFile(filePath string) (string, error) {
    sqlBytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        return "", fmt.Errorf("failed to read SQL file: %w", err)
    }
    return string(sqlBytes), nil
}