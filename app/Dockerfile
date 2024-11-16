# Golangイメージをベースにする
FROM golang:1.23.3-alpine

# 必要なツールをインストール
RUN apk add --no-cache git

# 作業ディレクトリを作成
WORKDIR /app

# go.modとgo.sumをコピーして依存関係を解決
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションコードをコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main .

# アプリケーションを起動
CMD ["./main"]
