# バックエンドログ機能 - 実装計画

## 1. 概要

Go標準ライブラリの `log/slog`（Go 1.21+）を使用してログ機能を実装する。

---

## 2. 実装フェーズ

### Phase 1: ロガーパッケージ作成

**ファイル**: `go/pkg/logger/logger.go`

```go
// 実装内容
- slogを使用したロガー初期化
- ログレベル設定（環境変数から読み込み）
- JSON/テキストフォーマット切り替え
- グローバルロガーインスタンス
```

**機能**:
- `Init()` - ロガーの初期化
- `Debug()`, `Info()`, `Warn()`, `Error()` - 各レベルのログ出力
- `WithRequestID()` - リクエストIDをコンテキストに追加

---

### Phase 2: ミドルウェア作成

**ファイル**: `go/middleware/logging.go`

```go
// 実装内容
- リクエストIDの生成/取得
- リクエスト開始時のログ出力
- レスポンス完了時のログ出力（ステータス、処理時間）
- コンテキストへのリクエストID格納
```

**機能**:
- `RequestLogger()` - Ginミドルウェア関数
- リクエストヘッダー `X-Request-ID` の処理
- レスポンスヘッダーへの `X-Request-ID` 付与

---

### Phase 3: ハンドラー・サービスへの組み込み

**ファイル**: `go/handlers/shuffle.go`（更新）

```go
// 追加内容
- バリデーションエラー時のログ出力
- シャッフル操作の開始/完了ログ
- エラー発生時の詳細ログ
```

**ファイル**: `go/services/shuffle.go`（更新）

```go
// 追加内容
- シャッフル処理の詳細ログ（DEBUGレベル）
- 結果サマリーのログ（INFOレベル）
```

---

### Phase 4: main.go 更新

**ファイル**: `go/main.go`（更新）

```go
// 追加内容
- ロガー初期化の呼び出し
- ロギングミドルウェアの登録
- Ginのデフォルトロガーを置き換え
```

---

## 3. ディレクトリ構造

```
go/
├── main.go                 (更新)
├── go.mod                  (更新: uuid パッケージ追加)
├── pkg/
│   └── logger/
│       └── logger.go       (新規)
├── middleware/
│   └── logging.go          (新規)
├── handlers/
│   └── shuffle.go          (更新)
├── services/
│   └── shuffle.go          (更新)
└── models/
    └── shuffle.go          (変更なし)
```

---

## 4. 依存関係

### 追加パッケージ
```
github.com/google/uuid  # リクエストID生成用
```

### 使用する標準ライブラリ
```
log/slog    # 構造化ログ
os          # 環境変数
time        # タイムスタンプ
```

---

## 5. 変更対象ファイル一覧

| ファイル | 操作 | 説明 |
|---------|------|------|
| `go/pkg/logger/logger.go` | 新規作成 | ロガーパッケージ |
| `go/middleware/logging.go` | 新規作成 | ログミドルウェア |
| `go/main.go` | 更新 | 初期化・ミドルウェア登録 |
| `go/handlers/shuffle.go` | 更新 | ハンドラーログ追加 |
| `go/services/shuffle.go` | 更新 | サービスログ追加 |
| `go/go.mod` | 更新 | uuid パッケージ追加 |

---

## 6. 実装順序

1. `go/pkg/logger/logger.go` - ロガーパッケージ作成
2. `go/middleware/logging.go` - ミドルウェア作成
3. `go/main.go` - 初期化とミドルウェア登録
4. `go/handlers/shuffle.go` - ハンドラーにログ追加
5. `go/services/shuffle.go` - サービスにログ追加
6. テスト実行・動作確認

---

## 7. テスト計画

### ユニットテスト
- `pkg/logger/logger_test.go` - ロガー初期化テスト
- `middleware/logging_test.go` - ミドルウェアテスト

### 統合テスト
- 既存の `handlers/shuffle_test.go` でログ出力を確認

### 手動テスト
- 各ログレベルでの出力確認
- JSON/テキスト形式の切り替え確認
- リクエストIDの伝播確認
