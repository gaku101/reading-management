# my-portfolio
作成中の読書管理アプリのバックエンド用リポジトリになります。  
フロントエンドのコードはhttps://github.com/gaku101/my-portfolio-front からご確認頂けます。  
近日中にこちらへアプリのURLを記載いたします。

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
│   └─ workflows
├─ api
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
