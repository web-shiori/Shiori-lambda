# Shiori-lambda
## 概要
- Web Snapshot(旧Web Shiori)のlambda処理
- S3に保存した画面のスクリーンショット(PDFをChromeで閲覧している状態を想定)から
「PDFの何ページ目を閲覧中か」という情報をOCRを利用して抽出し、contentテーブルの「pdf_page_num」カラムに保存する

#### 構成図
![構成図](./docs/PDF実装全体図.drawio.svg)

## 使用技術
- AWS Lambda
- S3
- Amazon Textract
- CloudWatch 
- Golang

## 開発手順
1. ビルド
```shell
make build
```

2. zipアップロード
