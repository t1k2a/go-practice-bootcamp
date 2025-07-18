# Go Practice Bootcamp 🧠💻

Go初心者でも、毎日少しずつ手を動かして楽しく学べる実践的カリキュラム。  
PHPやLaravel経験者がGoにスムーズに移行できるように設計しています。

---

## 📚 このリポジトリで学べること

- Goの基本文法（struct, error, method, goroutine, interface など）
- CLIツール開発、ファイルI/O、JSON/YAMLの扱い
- エラーハンドリングとユーティリティ関数の設計
- Goらしいアプリ設計思想とLaravelとの違いの理解

---

## 🧭 カリキュラム構成

| フェーズ | 期間 | 内容 |
|--------|------|------|
| Week1  | 2〜3週間 | CLI開発中心でGo文法に慣れる |
| Week2  | 2〜3週間 | REST API構築を通じてGoのWeb開発を学ぶ |

---

## 🗂 WEEK1：Go文法とCLI開発の基礎x

| Day | タイトル | 内容 |
|-----|----------|------|
| Day1 | Hello CLI | flagパッケージ入門／引数で挨拶 |
| Day2 | Text Analyzer | 標準入力・文字数/単語数の解析 |
| Day3 | Struct vs Class | Laravelとの違いを構造体で体験 |
| Day4 | Error Drill | カスタムエラー型・型アサーション |
| Day5 | Config Loader | JSON/YAMLファイルの読み込み |
| Day6 | ToDo CLI | 保存・更新・完了機能付きのCLI開発 |

---

---

## 🚀 WEEK2：REST API開発編

Goの標準ライブラリと軽量ルータを使って、実務レベルのREST APIを構築します。  
LaravelでのWeb開発経験がある人が「GoでAPIを作る感覚」にスムーズに慣れる設計です。

### 🔧 学ぶこと

- `net/http`, `chi`や`gorilla/mux`によるルーティング
- JSONエンコード／デコードとレスポンス処理
- バリデーション、構造体の役割分離（DTO, Handler, Service風）
- 簡易なRouter構成とMiddlewares
- 単体テスト・ハンドラテスト（`httptest`）

### 📅 想定カリキュラム（Day例）

| Day | 内容 | ファイル例 |
|-----|------|------------|
| Day1 | シンプルなGET API（/ping） | `/week2/day1-ping/` |
| Day2 | RESTful Todo API（GET/POST） | `/week2/day2-todo-api/` |
| Day3 | PUT/DELETEとバリデーション導入 | `/week2/day3-update-delete/` |
| Day4 | ミドルウェアでロギング＆ヘッダー | `/week2/day4-middleware/` |
| Day5 | サービス層・Handler層への分離 | `/week2/day5-structure/` |
| Day6 | テスト＆Curl/Postman操作体験 | `/week2/day6-testing/` |

---

## ☁️ WEEK3：デプロイ・実用構築編

CLIやAPIで作った成果物を外部に公開したり、よりプロダクション寄りの設計に近づけていきます。

### 🔧 学ぶこと

- Goアプリの環境変数管理（`.env` / `os.Getenv`)
- デプロイ先の選定（Render, Fly.io, Railwayなど）
- Docker化してビルド・起動（簡易でOK）
- GitHub ActionsによるCI/CDの導入
- READMEやバッジの整備、ポートフォリオ化

### 📅 想定カリキュラム（Day例）

| Day | 内容 | フォルダ例 |
|-----|------|-----------|
| Day1 | `.env`対応と環境別設定 | `/week3/day1-env/` |
| Day2 | Render/Fly.ioへ初デプロイ | `/week3/day2-deploy/` |
| Day3 | Docker対応とローカル実行 | `/week3/day3-docker/` |
| Day4 | CI/CD入門（GitHub Actions） | `/week3/day4-ci-cd/` |
| Day5 | README仕上げ・リリース体験 | `/week3/day5-readme/` |
| Day6 | 公開ポートフォリオとして整備 | `/week3/day6-finish/` |

---

## 🗒 補足

WEEK2, WEEK3は現在開発中のロードマップです。  
IssueやPR、提案など大歓迎です！

---

## 🤖 AI支援学習におすすめのプロンプト

ChatGPTやClaudeなどに以下のプロンプトを投げると、Day単位で一緒に進めてくれます。
適宜変更してください。

```text
あなたはGo言語教育者兼テックリードです。
私は${経験年数、得意言語またはFW}のエンジニアで、Goを学びたいです。
Go言語基礎 → CLI／API／DB／並行処理などの小タスク学習 → CLI ToDoアプリ → Week2 API開発、と続けるカリキュラムです。
毎日30分、退屈しない実装中心で教えてください。
ユーザーが「DayNやりたい」「このコードを評価して」「次進めたい」と言ったら、即コード提示＋レビューで導いてください。
```

