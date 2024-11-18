# polaris-api

## Build & Deploy 手順

Github Actions を使った CI/CD を実現するまでは手動でビルドし、デプロイする。

### Build

プロジェクトの app ディレクトリに移動し、ビルドコマンドを実行する

```sh
# appディレクトリに移動
$ cd app

# ビルド（main.goをbootstrapというアーティファクトにする。名前厳守。実行環境はLambdaのarm64）
$ GOARCH=arm64 GOOS=linux go build -tags lambda.norpc -o bootstrap main.go
```

Lambda へ手動デプロイするために、ビルドしたアーティファクトを zip 化する

```sh
# ビルドしたアーティファクト（bootstrap）をdeployment.zipにする（ファイル名は何でもいい）
zip deployment.zip bootstrap
```

### Deploy

deployment.zip を AWS Lambda に手動デプロイする

1. [Lambda](https://ap-northeast-1.console.aws.amazon.com/lambda/home?region=ap-northeast-1#/functions/polaris-api)にアクセスする
2. コードソースの「アップロード元」ボタンを押下し、「.zip ファイル」を選択する
3. 「アップロード」ボタンで`deployment.zip`を選択し、保存する
