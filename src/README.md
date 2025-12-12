# オフィス出社状況管理システム

オフィスへの出社状況を管理するためのRailsアプリケーションです。従業員の名前と位置情報を登録・管理できます。

## 機能

- 出社状況の登録（名前と位置）
- 出社状況の一覧表示
- 出社状況の編集
- 出社状況の削除

## 技術スタック

- **Ruby**: 3.2.7
- **Rails**: 7.1.5
- **データベース**: MySQL 8.0
- **Webサーバー**: Puma
- **フロントエンド**: Hotwire (Turbo + Stimulus)

## セットアップ

### 必要な環境

- Docker と Docker Compose
- または Ruby 3.2.7 と MySQL 8.0

### Dockerを使用する場合

1. リポジトリをクローンします：
```bash
cd src
```

2. Docker Composeでアプリケーションを起動します：
```bash
docker-compose up --build
```

3. データベースをセットアップします（別のターミナルで）：
```bash
docker-compose exec web rails db:create db:migrate
```

4. ブラウザで `http://localhost:3001` にアクセスします。

### ローカル環境で実行する場合

1. 依存関係をインストールします：
```bash
bundle install
```

2. データベースを作成・マイグレートします：
```bash
rails db:create
rails db:migrate
```

3. サーバーを起動します：
```bash
rails server
```

4. ブラウザで `http://localhost:3000` にアクセスします。

## データベース

### テーブル構造

**attendances**
- `id` (integer, primary key)
- `name` (string) - 従業員の名前
- `location` (string) - 位置情報
- `created_at` (datetime)
- `updated_at` (datetime)

### バリデーション

- `name`: 必須
- `location`: 必須

## ルーティング

- `GET /` - 出社状況一覧
- `GET /attendances/new` - 新規登録フォーム
- `POST /attendances` - 出社状況の作成
- `GET /attendances/:id/edit` - 編集フォーム
- `PATCH /attendances/:id` - 出社状況の更新
- `DELETE /attendances/:id` - 出社状況の削除

## Docker環境

### サービス構成

- **web**: Railsアプリケーション（ポート3001）
- **db**: MySQL 8.0（ポート3306）

### 環境変数

- `RAILS_ENV`: development
- `DATABASE_HOST`: db
- `DATABASE_USERNAME`: user
- `DATABASE_PASSWORD`: password
- `DATABASE_NAME`: hackathon_b_team_development

## 開発

### データベースのリセット

```bash
# Docker環境の場合
docker-compose exec web rails db:reset

# ローカル環境の場合
rails db:reset
```

### コンソールの起動

```bash
# Docker環境の場合
docker-compose exec web rails console

# ローカル環境の場合
rails console
```

## ライセンス

このプロジェクトはハッカソン用のアプリケーションです。
