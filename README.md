# Go Todo API

Go言語の学習を目的として作成した、シンプルなTodo管理のためのREST APIアプリケーションです。
WebフレームワークにGin、ORMにGORMを採用し、PostgreSQLと連携してCRUD（作成・読み取り・更新・削除）処理を実装しています。

## 🛠 テクノロジースタック

* **言語**: Go
* **Webフレームワーク**: Gin
* **ORM**: GORM
* **データベース**: PostgreSQL
* **環境変数管理**: godotenv

## ディレクトリ構成

\`\`\`text
01_go/
├── main.go                # アプリケーションのエントリーポイント・ルーティング設定
├── go.mod                 # パッケージ管理ファイル
├── go.sum                 # パッケージの依存関係バージョン情報
├── .env                   # データベース接続情報など（※Git管理外）
├── database/
│   └── database.go        # PostgreSQLへの接続処理・マイグレーション
├── models/
│   └── todo.go            # Todoのデータ構造（GORMモデル）
└── controllers/
    └── todo_controller.go # 各APIエンドポイントのビジネスロジック (CRUD処理)

##  環境構築と起動方法

### 1. 前提条件
* Go がインストールされていること
* PostgreSQL がインストールされ、起動していること

### 2. データベースの準備
PostgreSQLにアクセスし、本アプリ用のデータベースを作成します。
```
sudo -u postgres psql
sql
CREATE DATABASE todo_db;
\q
```

### 3. 環境変数の設定
プロジェクトのルートディレクトリ（`main.go` と同じ階層）に `.env` ファイルを作成し、以下の内容を記述します。
※環境に合わせてユーザー名やパスワードを変更してください。

```
env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_DATABASE=todo_db
```

### 4. パッケージのインストール
依存関係をダウンロードします。

```
go mod tidy
```

### 5. アプリケーションの起動
サーバーを起動します。起動時に自動的にデータベースのマイグレーション（テーブル作成）が行われます。

```
go run main.go
```
起動後、http://localhost:8080 でサーバーが待機状態になります。

## API エンドポイント一覧

| メソッド | エンドポイント | 説明 | リクエストボディ例 |
| :--- | :--- | :--- | :--- |
| **GET** | `/todos` | Todoの一覧を取得します | - |
| **POST** | `/todos` | 新しいTodoを作成します | `{"title": "Goの学習", "memo": "Ginを使う"}` |
| **PUT** | `/todos/:id` | 指定したIDのTodoを更新します | `{"status": "completed"}` |
| **DELETE** | `/todos/:id` | 指定したIDのTodoを削除します | - |

### 動作確認（curlコマンドの例）

**Todoの作成**

```
curl -X POST http://localhost:8080/todos \
-H "Content-Type: application/json" \
-d '{"title": "Go言語の学習をすすめる", "memo": "APIを完成させる"}'
```

**Todoの一覧取得**
```
curl http://localhost:8080/todos
```