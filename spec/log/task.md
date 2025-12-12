# バックエンドログ機能 - タスク一覧

## 概要

実装計画を元に切り出したタスク一覧。チェックボックスで進捗管理を行う。

---

## Phase 1: ロガーパッケージ作成

### 1.1 ディレクトリ作成
- [ ] `go/pkg/logger/` ディレクトリを作成

### 1.2 ロガー実装 (`go/pkg/logger/logger.go`)
- [ ] パッケージ定義と import 文の記述
- [ ] ログレベル定数の定義（DEBUG, INFO, WARN, ERROR）
- [ ] 環境変数 `LOG_LEVEL` からログレベルを読み込む関数
- [ ] 環境変数 `LOG_FORMAT` からフォーマットを読み込む関数
- [ ] `Init()` 関数の実装（slog ハンドラー初期化）
- [ ] JSON ハンドラーの設定
- [ ] テキストハンドラーの設定
- [ ] `WithRequestID()` 関数の実装

### 1.3 ロガーテスト (`go/pkg/logger/logger_test.go`)
- [ ] `Init()` 関数のテスト
- [ ] ログレベル設定のテスト
- [ ] フォーマット切り替えのテスト

---

## Phase 2: ミドルウェア作成

### 2.1 ディレクトリ作成
- [ ] `go/middleware/` ディレクトリを作成

### 2.2 ミドルウェア実装 (`go/middleware/logging.go`)
- [ ] パッケージ定義と import 文の記述
- [ ] リクエストID生成関数（UUID v4）
- [ ] `X-Request-ID` ヘッダーからのID取得処理
- [ ] レスポンスヘッダーへの `X-Request-ID` 付与
- [ ] リクエスト開始時のログ出力
- [ ] レスポンス完了時のログ出力（ステータス、処理時間）
- [ ] `RequestLogger()` Gin ミドルウェア関数の実装
- [ ] コンテキストへのリクエストID格納

### 2.3 ミドルウェアテスト (`go/middleware/logging_test.go`)
- [ ] リクエストID生成のテスト
- [ ] ヘッダーからのID取得テスト
- [ ] ログ出力のテスト

---

## Phase 3: 依存関係の更新

### 3.1 go.mod 更新
- [ ] `github.com/google/uuid` パッケージを追加
- [ ] `go mod tidy` を実行

---

## Phase 4: main.go 更新

### 4.1 ロガー初期化
- [ ] logger パッケージの import 追加
- [ ] middleware パッケージの import 追加
- [ ] `main()` 関数内で `logger.Init()` を呼び出し

### 4.2 ミドルウェア登録
- [ ] Gin のデフォルトロガーを無効化（`gin.New()` に変更）
- [ ] `middleware.RequestLogger()` を登録
- [ ] Recovery ミドルウェアの追加

---

## Phase 5: ハンドラーへのログ追加

### 5.1 handlers/shuffle.go 更新
- [ ] logger パッケージの import 追加
- [ ] バリデーションエラー時の WARN ログ追加
- [ ] シャッフル操作開始時の INFO ログ追加
- [ ] シャッフル操作完了時の INFO ログ追加（結果サマリー）
- [ ] エラー発生時の ERROR ログ追加

---

## Phase 6: サービスへのログ追加

### 6.1 services/shuffle.go 更新
- [ ] logger パッケージの import 追加
- [ ] シャッフル処理開始時の DEBUG ログ追加
- [ ] グループ分割処理の DEBUG ログ追加
- [ ] 処理完了時の INFO ログ追加

---

## Phase 7: テスト・動作確認

### 7.1 ユニットテスト
- [ ] `go test ./pkg/logger/...` を実行
- [ ] `go test ./middleware/...` を実行

### 7.2 既存テストの確認
- [ ] `go test ./handlers/...` を実行
- [ ] `go test ./services/...` を実行

### 7.3 統合テスト
- [ ] `go test ./...` を実行して全テストパス確認

### 7.4 手動テスト
- [ ] `LOG_LEVEL=DEBUG` でのログ出力確認
- [ ] `LOG_LEVEL=INFO` でのログ出力確認
- [ ] `LOG_FORMAT=json` でのフォーマット確認
- [ ] `LOG_FORMAT=text` でのフォーマット確認
- [ ] リクエストIDの伝播確認（`X-Request-ID` ヘッダー）

### 7.5 Docker 環境テスト
- [ ] `docker compose up --build` でビルド確認
- [ ] コンテナ内でのログ出力確認

---

## タスクサマリー

| Phase | タスク数 | 説明 |
|-------|---------|------|
| Phase 1 | 11 | ロガーパッケージ作成 |
| Phase 2 | 11 | ミドルウェア作成 |
| Phase 3 | 2 | 依存関係の更新 |
| Phase 4 | 5 | main.go 更新 |
| Phase 5 | 5 | ハンドラーへのログ追加 |
| Phase 6 | 4 | サービスへのログ追加 |
| Phase 7 | 10 | テスト・動作確認 |
| **合計** | **48** | |

---

## 優先度

1. **高**: Phase 1, 2, 3, 4（基盤となる機能）
2. **中**: Phase 5, 6（既存コードへの組み込み）
3. **低**: Phase 7（テスト・確認）
