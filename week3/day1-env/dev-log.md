# Week3 Day1 - 環境変数の管理

## 学習内容

### godotenvパッケージの使用
- `github.com/joho/godotenv`を使って`.env`ファイルから環境変数を読み込む方法を学習
- `godotenv.Load()`で`.env`ファイルを読み込み、`os.Getenv()`で値を取得

### 環境別設定の切り替え
- `APP_ENV`環境変数を使って開発環境（dev）と本番環境（prod）を判別
- 開発環境では`.env.local`ファイルを追加で読み込むように実装

### 遭遇した問題と解決

#### 1. .env.localの値が反映されない問題
**問題**: `.env`と`.env.local`の両方に同じ環境変数が定義されている場合、`.env.local`の値が反映されなかった

**原因**: `godotenv.Load()`は既存の環境変数を上書きしないため、先に読み込まれた`.env`の値が優先される

**解決**: `godotenv.Overload()`を使用することで、既存の環境変数を上書きして`.env.local`の値を反映

```go
// 変更前
godotenv.Load(".env.local")

// 変更後
godotenv.Overload(".env.local")
```

#### 2. ファイル存在チェックの条件ミス
**問題**: `os.Stat()`のエラーチェックで`err != nil`としていたため、ファイルが存在しない時に読み込もうとしていた

**解決**: `err == nil`に修正（エラーがない = ファイルが存在する）

### 実装のポイント
- 環境変数の読み込み順序が重要
- オプショナルなファイル（.env.local）は存在チェックしてから読み込む
- 開発環境と本番環境で異なる設定を使い分けることで、環境ごとの設定管理が容易に

### 動作確認
- `.env`: `PORT=8081`, `APP_NAME=GoEnvApp`, `APP_ENV=prod`
- `.env.local`: `PORT=8082`, `APP_NAME=GoEnvAppLocal`, `APP_ENV=dev`
- 開発環境（APP_ENV=dev）では`.env.local`の値が優先されることを確認