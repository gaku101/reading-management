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
- プロフィール画像の変更

＊現在その他機能追加やレスポンシブ対応などを行なっています。


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
- sqlc(ORM)
- viper
- RDS(PostgreSQL)
- Github Actions
- Docker 
- Kubernetes
- S3
- ECR
- EKS
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
## ローカル開発環境のセットアップ

### ツールのインストール

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install sqlc
    ```

- [Gomock](https://github.com/golang/mock)

    ``` bash
    go install github.com/golang/mock/mockgen@v1.6.0
    ```

### インフラのセットアップ

- Create the bank-network

    ``` bash
    make network
    ```

- Start postgres container:

    ```bash
    make postgres
    ```

- Create simple_bank database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Run db migration up 1 version:

    ```bash
    make migrateup1
    ```

- Run db migration down all versions:

    ```bash
    make migratedown
    ```

- Run db migration down 1 version:

    ```bash
    make migratedown1
    ```

### コードの生成方法

- Generate SQL CRUD with sqlc:

    ```bash
    make sqlc
    ```

- Generate DB mock with gomock:

    ```bash
    make mock
    ```

- Create a new db migration:

    ```bash
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```

### コードの実行

- Run server:

    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```

## kubernetes clusterへのデプロイ

- [Install nginx ingress controller](https://kubernetes.github.io/ingress-nginx/deploy/#aws):

    ```bash
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.48.1/deploy/static/provider/aws/deploy.yaml
    ```

- [Install cert-manager](https://cert-manager.io/docs/installation/kubernetes/):

    ```bash
    kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml
    ```
