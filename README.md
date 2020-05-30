# backend

[![reviewdog](https://github.com/wheatandcat/PeperomiaBackend/workflows/reviewdog/badge.svg?branch=master&event=push)](https://github.com/wheatandcat/PeperomiaBackend/actions?query=workflow%3Areviewdog+event%3Apush+branch%3Amaster)

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
$ dev_appserver.py app.yaml
```

## テスト

```
$ go test ./handler
```


カバレッジ確認

```
$ go test -coverprofile=cover.out ./handler
```

## デプロイ

```
$ gcloud app deploy
```

# ツール

## APIドキュメント

https://app.swaggerhub.com/apis-docs/wheatandcat/peperomia/1.0.0