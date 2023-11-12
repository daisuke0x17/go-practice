## 資料
- https://qiita.com/ta1m1kam/items/fc798cdd6a4eaf9a7d5e

## Goroutines
- goroutineは軽量のスレッド
    - 並列実行される

## Channels
- 並列実行されるgoroutine間を接続するパイプ（トンネル）
    - 並列実行している関数から値を取得する
    - あるgoroutineから別のgoroutineへ値を渡す
- チャンネルから値を受け取るのが他の非同期処理扱える言語の`await`に近い部分