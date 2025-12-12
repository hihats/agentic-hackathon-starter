# シャッフルランチアプリ - 実装計画

## 概要
参加者をランダムにシャッフルしてグループ分けするWebアプリケーション

## 技術スタック
- **フロントエンド**: Nuxt 3 / TypeScript (port 3000)
- **バックエンド**: Go / Gin (port 8080)

---

## Phase 1: Goバックエンド

### 1.1 ディレクトリ構造の作成
```
go/
├── main.go          (更新)
├── models/
│   └── shuffle.go   (新規)
├── services/
│   └── shuffle.go   (新規)
└── handlers/
    └── shuffle.go   (新規)
```

### 1.2 モデル定義
**ファイル**: `go/models/shuffle.go`
- `ShuffleRequest`: participants ([]string), group_size (*int), num_groups (*int)
- `ShuffleResponse`: groups ([][]string)
- `ErrorResponse`: error, message

### 1.3 シャッフルサービス
**ファイル**: `go/services/shuffle.go`
- ランダムシャッフル処理
- グループ分割ロジック
- 余りメンバーの均等配分

### 1.4 HTTPハンドラー
**ファイル**: `go/handlers/shuffle.go`
- JSONバインディング
- バリデーション (4名以上、2グループ以上)
- エラーレスポンス

### 1.5 main.goの更新
**ファイル**: `go/main.go`
- CORSミドルウェア追加
- `POST /api/shuffle` エンドポイント追加

---

## Phase 2: Nuxtフロントエンド

### 2.1 ディレクトリ構造
```
nuxt/
├── types/
│   └── shuffle.ts       (新規)
├── composables/
│   └── useShuffleApi.ts (新規)
├── components/
│   ├── ParticipantInput.vue  (新規)
│   ├── GroupSettings.vue     (新規)
│   └── ResultDisplay.vue     (新規)
└── pages/
    └── index.vue        (更新)
```

### 2.2 型定義
**ファイル**: `nuxt/types/shuffle.ts`
- ShuffleRequest, ShuffleResponse, ErrorResponse インターフェース

### 2.3 APIクライアント
**ファイル**: `nuxt/composables/useShuffleApi.ts`
- fetch を使用したAPI呼び出し
- エラーハンドリング

### 2.4 コンポーネント

**ParticipantInput.vue**
- テキストエリアで名前入力
- 改行/カンマ区切り対応
- 参加者数カウント表示

**GroupSettings.vue**
- グループサイズ or グループ数の選択
- ラジオボタンで切り替え

**ResultDisplay.vue**
- グループごとの結果表示
- グループ名 (A, B, C...)

### 2.5 メインページ更新
**ファイル**: `nuxt/pages/index.vue`
- コンポーネント統合
- シャッフルボタン
- ローディング/エラー表示

---

## バリデーションルール

| ルール | 値 |
|--------|-----|
| 最小参加者数 | 4名 |
| 最小グループ数 | 2グループ |
| group_size | 1以上 |
| num_groups | 2以上 |

---

## APIエンドポイント

### POST /api/shuffle

**リクエスト:**
```json
{
  "participants": ["Alice", "Bob", "Charlie", "Dave"],
  "group_size": 2
}
```

**レスポンス:**
```json
{
  "groups": [
    ["Charlie", "Alice"],
    ["Bob", "Dave"]
  ]
}
```

---

## 実装順序

1. Go: models/shuffle.go
2. Go: services/shuffle.go
3. Go: handlers/shuffle.go
4. Go: main.go (CORS + ルート)
5. Nuxt: types/shuffle.ts
6. Nuxt: composables/useShuffleApi.ts
7. Nuxt: components (3ファイル)
8. Nuxt: pages/index.vue

---

## 変更対象ファイル一覧

| ファイル | 操作 |
|---------|------|
| `go/models/shuffle.go` | 新規作成 |
| `go/services/shuffle.go` | 新規作成 |
| `go/handlers/shuffle.go` | 新規作成 |
| `go/main.go` | 更新 |
| `nuxt/types/shuffle.ts` | 新規作成 |
| `nuxt/composables/useShuffleApi.ts` | 新規作成 |
| `nuxt/components/ParticipantInput.vue` | 新規作成 |
| `nuxt/components/GroupSettings.vue` | 新規作成 |
| `nuxt/components/ResultDisplay.vue` | 新規作成 |
| `nuxt/pages/index.vue` | 更新 |
