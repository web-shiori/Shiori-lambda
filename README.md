# Shiori-lambda
## 概要
- Web Snapshot(旧Web Shiori)のlambda処理
- S3に保存した画面のスクリーンショット(PDFをChromeで閲覧している状態を想定)から
「PDFの何ページ目を閲覧中か」という情報をOCRを利用して抽出し、contentテーブルの「」カラムに保存する
## 使用技術
- AWS Lambda
- S3
- Amazon Textract
- CloudWatch
- Serveless Framework
- Golang

## 開発手順
### ビルド
## TODO: 修正
```shell
% GOOS=linux go build -o lambda-page-number-extract
```
## デプロイ
```shell
# zip作成
% zip lambda.zip Shiori-lambda
```

- zipアップロード
