# 機能追加仕様書: お店リストの画像付きUI表示

## 概要

食べログから取得したお店リストを、テキストエリアではなく画像付きのカード形式UIで表示する。ユーザーはカードをクリックして店舗を選択/解除でき、選択した店舗のみがシャッフル対象となる。

## 要件

### 機能概要

1. **ページ初期表示時に食べログから店舗情報を自動取得**（デフォルト定数を廃止）
2. 取得した店舗情報を画像付きカードで**横スクロール**表示
3. カードクリックで選択/解除（トグル）
4. 選択された店舗のみがシャッフル対象
5. 全選択/全解除/再取得ボタン

### 取得する追加情報

| 項目 | 用途 |
|------|------|
| 店舗画像URL | カードに表示 |
| 評価スコア（任意） | カードに表示 |
| ジャンル | カードに表示 |

## UI設計

### 変更前（テキストエリア）

```
+------------------------------------------+
| お店リスト（1行に1店舗）                  |
| [周辺のお店を取得]                        |
| +--------------------------------------+ |
| | 中華料理店                           | |
| | イタリアン                           | |
| | ...                                  | |
| +--------------------------------------+ |
+------------------------------------------+
```

### 変更後（画像付きカード・横スクロール）

```
+------------------------------------------------------------------+
| お店リスト                                                        |
| [再取得] [全選択] [全解除]                                        |
+------------------------------------------------------------------+
| ← +--------+  +--------+  +--------+  +--------+  +--------+ →   |
|   |  画像  |  |  画像  |  |  画像  |  |  画像  |  |  画像  |      |
|   +--------+  +--------+  +--------+  +--------+  +--------+      |
|   | 店名   |  | 店名   |  | 店名   |  | 店名   |  | 店名   |      |
|   | ジャンル|  | ジャンル|  | ジャンル|  | ジャンル|  | ジャンル|      |
|   | ★ 3.5 |  | ★ 3.8 |  | ★ 4.0 |  | ★ 3.2 |  | ★ 3.9 |      |
|   | [✓選択]|  | [  ]   |  | [✓選択]|  | [✓選択]|  | [  ]   |      |
|   +--------+  +--------+  +--------+  +--------+  +--------+      |
+------------------------------------------------------------------+
                        ← 横スクロール可能 →
```

### カード仕様

- **サイズ**: 幅150px程度、レスポンシブ対応
- **画像**: 正方形またはアスペクト比4:3、上部に配置
- **選択状態**:
  - 未選択: 半透明、グレーアウト
  - 選択済: 通常表示、チェックマーク、ボーダーハイライト
- **ホバー**: スケールアップ、シャドウ強調

## 技術設計

### バックエンドAPI変更

**エンドポイント**: `GET /api/restaurants/nearby`（変更なし）

**レスポンス例（拡張）**:
```json
{
  "restaurants": [
    {
      "name": "イタリアン トラットリア",
      "genre": "イタリアン",
      "imageUrl": "https://tblg.k-img.com/restaurant/images/...",
      "rating": 3.52,
      "url": "https://tabelog.com/tokyo/A1307/..."
    }
  ]
}
```

### スクレイピングサービス変更

**ファイル**: `app/services/tabelog_scraper_service.rb`

**追加取得項目**:
```ruby
{
  name: "店舗名",
  genre: "ジャンル",
  image_url: "画像URL",
  rating: 3.52,
  url: "店舗詳細URL"
}
```

### フロントエンド型定義変更

**ファイル**: `frontend/src/lib/api.ts`

```typescript
export interface Restaurant {
  name: string;
  genre: string;
  imageUrl: string;
  rating: number | null;
  url: string;
}
```

### フロントエンドUI変更

**ファイル**: `frontend/src/routes/+page.svelte`

**状態管理**:
```typescript
import { onMount } from 'svelte';

let fetchedRestaurants: Restaurant[] = [];
let selectedRestaurants: Set<string> = new Set();
let isLoadingRestaurants = false;

// 初期表示時に自動で食べログから取得
onMount(async () => {
  await loadRestaurants();
});

async function loadRestaurants() {
  isLoadingRestaurants = true;
  try {
    fetchedRestaurants = await restaurantService.fetchNearby();
    // 取得後は全店舗を選択状態にする
    selectedRestaurants = new Set(fetchedRestaurants.map(r => r.name));
  } catch (err) {
    console.error('店舗情報の取得に失敗しました', err);
  } finally {
    isLoadingRestaurants = false;
  }
}

function toggleRestaurant(name: string) {
  if (selectedRestaurants.has(name)) {
    selectedRestaurants.delete(name);
  } else {
    selectedRestaurants.add(name);
  }
  selectedRestaurants = selectedRestaurants; // リアクティブ更新
}

function selectAll() {
  selectedRestaurants = new Set(fetchedRestaurants.map(r => r.name));
}

function deselectAll() {
  selectedRestaurants = new Set();
}
```

**デフォルト定数の削除**:
- `let restaurants = '中華料理店\nイタリアン\n...'` を削除
- 代わりに `fetchedRestaurants` と `selectedRestaurants` を使用

**シャッフル時の店舗リスト取得**:
```typescript
const restaurantList = Array.from(selectedRestaurants);
```

### コンポーネント構造

```svelte
<div class="restaurant-section">
  <div class="restaurant-header">
    <h2>お店リスト</h2>
    <div class="restaurant-actions">
      <button on:click={loadRestaurants} disabled={isLoadingRestaurants}>
        {isLoadingRestaurants ? '取得中...' : '再取得'}
      </button>
      <button on:click={selectAll}>全選択</button>
      <button on:click={deselectAll}>全解除</button>
    </div>
  </div>

  {#if isLoadingRestaurants}
    <div class="loading">店舗情報を取得中...</div>
  {:else if fetchedRestaurants.length > 0}
    <div class="restaurant-scroll-container">
      {#each fetchedRestaurants as restaurant}
        <div
          class="restaurant-card"
          class:selected={selectedRestaurants.has(restaurant.name)}
          on:click={() => toggleRestaurant(restaurant.name)}
        >
          <img src={restaurant.imageUrl} alt={restaurant.name} loading="lazy" />
          <div class="restaurant-info">
            <h3>{restaurant.name}</h3>
            <span class="genre">{restaurant.genre}</span>
            {#if restaurant.rating}
              <span class="rating">★ {restaurant.rating}</span>
            {/if}
          </div>
          <div class="checkbox">
            {selectedRestaurants.has(restaurant.name) ? '✓' : ''}
          </div>
        </div>
      {/each}
    </div>
  {:else}
    <p class="placeholder">店舗情報を取得できませんでした</p>
  {/if}
</div>
```

## 実装ファイル

| ファイル | 操作 | 内容 |
|----------|------|------|
| `app/services/tabelog_scraper_service.rb` | 編集 | 画像URL・評価取得追加 |
| `app/controllers/api/restaurants_controller.rb` | 編集 | レスポンスキー変換（snake_case → camelCase） |
| `frontend/src/lib/api.ts` | 編集 | Restaurant型拡張 |
| `frontend/src/routes/+page.svelte` | 編集 | カードUI実装、選択機能追加 |

## CSSスタイル案

```css
.restaurant-section {
  margin-bottom: 24px;
}

.restaurant-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.restaurant-actions {
  display: flex;
  gap: 8px;
}

.restaurant-actions button {
  padding: 8px 16px;
  border-radius: 8px;
  font-size: 12px;
}

/* 横スクロールコンテナ */
.restaurant-scroll-container {
  display: flex;
  gap: 12px;
  overflow-x: auto;
  padding: 8px 0;
  scroll-behavior: smooth;
  -webkit-overflow-scrolling: touch;
}

/* スクロールバーのスタイリング */
.restaurant-scroll-container::-webkit-scrollbar {
  height: 8px;
}

.restaurant-scroll-container::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 4px;
}

.restaurant-scroll-container::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.3);
  border-radius: 4px;
}

.restaurant-card {
  flex: 0 0 140px;  /* 固定幅で横スクロール */
  background: rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.2s ease;
  opacity: 0.5;
  border: 2px solid transparent;
}

.restaurant-card:hover {
  transform: scale(1.02);
}

.restaurant-card.selected {
  opacity: 1;
  border: 2px solid #ffd93d;
  box-shadow: 0 4px 12px rgba(255, 217, 61, 0.3);
}

.restaurant-card img {
  width: 100%;
  aspect-ratio: 4/3;
  object-fit: cover;
}

.restaurant-info {
  padding: 8px;
  color: white;
}

.restaurant-info h3 {
  font-size: 11px;
  font-weight: 600;
  margin: 0 0 4px 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.genre {
  font-size: 10px;
  opacity: 0.7;
  display: block;
}

.rating {
  color: #ffd93d;
  font-size: 11px;
  font-weight: 600;
}

.checkbox {
  position: absolute;
  top: 8px;
  right: 8px;
  width: 20px;
  height: 20px;
  background: rgba(0, 0, 0, 0.5);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffd93d;
  font-size: 12px;
}

.loading {
  text-align: center;
  padding: 40px;
  color: rgba(255, 255, 255, 0.7);
}
```

## 注意事項

### 画像取得について

- 食べログの画像URLは外部参照可能か確認が必要
- 画像が取得できない場合はプレースホルダー画像を表示
- CORS対策が必要な場合はプロキシ経由を検討

### パフォーマンス

- 画像の遅延読み込み（lazy loading）を実装
- 画像サイズは小さめのサムネイルを使用

## テスト観点

- [ ] **初期表示時**に自動で店舗情報が取得されること
- [ ] 取得中はローディング表示されること
- [ ] 店舗取得後に画像付きカードが**横スクロール**で表示されること
- [ ] カードクリックで選択/解除が切り替わること
- [ ] 選択状態が視覚的に分かること（黄色ボーダー、チェックマーク）
- [ ] 全選択/全解除/再取得ボタンが機能すること
- [ ] 選択した店舗のみがシャッフル対象になること
- [ ] 画像が読み込めない場合にプレースホルダーが表示されること
- [ ] 横スクロールがスムーズに動作すること
