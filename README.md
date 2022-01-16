# Shiori-lambda
Web Snapshot(旧Web Shiori)のlambda関数

## 概要
- ./extract-pdf-page-num
- ./invert-image-color

### extract-pdf-page-num
保存されたPDFのスクリーンショットから現在のページ数を抽出し、データベースにPUTリクエストを送る。

### invert-image-color
S3に保存された画像の色を反転させる処理。

#### 構成図
![構成図](./docs/PDF実装全体図.drawio.svg)

## 使用技術
- AWS Lambda
- S3
- Amazon Textract
- CloudWatch 
- Golang

## 開発手順
1. zip作成
```shell
make zip
```

2. zipアップロード
