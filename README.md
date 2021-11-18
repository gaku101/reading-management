# my-portfolio(https://reading-management.net)
作成中の読書管理アプリのバックエンド用リポジトリになります。  
https://reading-management.net でデプロイしたアプリを実際にご利用頂けます(＊現在いくつかのバグを修正中です)。  
フロントエンドのコードはhttps://github.com/gaku101/my-portfolio-front からご確認頂けます。  

### 主な機能
- 読書記録の作成・閲覧・削除
- Google Books APIを利用した書籍情報の検索 & 読書記録作成
- 読書記録にノート(メモ)を追加・編集
- 読書記録に読んだページ数を登録
- 読書記録にカテゴリーを追加
- 読書記録にコメントを追加
- 読書記録をお気に入りに追加(⭐️ボタン押下)
- ユーザーのフォロー
- ユーザーのサインアップ・サインイン
- ユーザーの削除
- ポイントの獲得、他ユーザーへの送付
- 獲得ポイントに基づくランク制度(プロフィールで現在のランクを確認)
- プロフィール画像の変更(デプロイ環境でのバグを修正中)


## 検証用アカウント

id: testuser  
password: test01

---

## バージョン情報

- go 1.16-alpine3.13
- PostgreSQL 12-alpine
- golang-migrate v4.14.1

---

## 使用技術

- Go  
- gin
- sqlc
- viper
- PostgreSQL
- Github Actions
- Docker 
- Kubernetes
---


## ディレクトリ構成

```
@
├─ github
│   └─ workflows　// github-actionsの設定ファイル
├─ api // 各種API
├─ db
│   ├─ migration
│   ├─ mock
│   ├─ query
│   └─ sqlc
├─ infrastructure
├─ token
├─ util
├─ app.env
├─ docker-compose.yaml
├─ Dockerfile
├─ go.mod
├─ go.sum
├─ main.go
├─ Makefile
├─ sqlc.yaml
├─ start.sh
└─ wait-for.sh
```
