# backend

[![reviewdog](https://github.com/wheatandcat/PeperomiaBackend/workflows/reviewdog/badge.svg?branch=master&event=push)](https://github.com/wheatandcat/PeperomiaBackend/actions?query=workflow%3Areviewdog+event%3Apush+branch%3Amaster) [![codecov](https://codecov.io/gh/wheatandcat/PeperomiaBackend/branch/master/graph/badge.svg)](https://codecov.io/gh/wheatandcat/PeperomiaBackend)

## 準備

```
$ go mod download
```

or

```
$ go mod tidy
```


## ローカル実行

```
$ dev_appserver.py app.local.yaml
```

### GraphQL スキーマドキュメント

```
$ npx graphql-markdown http://localhost:8080/graphql > schema.md
```


## テスト

### mock作成

```
$ sh mock_gen.sh
```


### 実行

```
$ go test -race ./...
```


カバレッジ確認

```
$ go test -coverprofile=cover.out ./handler
$ go tool cover -html=cover.out -o cover.html
$ open cover.html
```

## GraphQLスキーマ更新

```
$ go run github.com/99designs/gqlgen generate
```

## デプロイ

```
$ gcloud app deploy
```

```
$ gcloud app deploy cron.yaml
```

# ツール

## APIドキュメント

https://app.swaggerhub.com/apis-docs/wheatandcat/peperomia/1.0.0
