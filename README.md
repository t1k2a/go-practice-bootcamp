# Go Practice Bootcamp 🧠💻

Go言語を体系的に・実践的に学ぶためのトレーニングリポジトリです。  
「文法だけで終わらせない」「退屈しない」「毎週動くものを作る」をテーマに、2〜3週間でGoの基本を習得します。

---

## 🎯 目的

- Go言語の基本文法・構造・ツール群を短期間で習得
- Laravel経験者が違和感なくGoに移行できるような比較と設計視点を養う
- 毎週小さなプロダクトを作りながら、Goらしい開発感覚を体験
- フェーズ2では実践的なポートフォリオを構築

---

## 🧩 カリキュラム構成

### 🔹 フェーズ1：基礎 × 実装トレーニング（2〜3週間）

| 週 | ミニゴール | 内容 |
|----|------------|------|
| Week1 | CLI開発に慣れる | `flag`パッケージ, struct, error handling |
| Week2 | 簡易REST API開発 | `net/http`, JSON, goroutine |
| Week3 | DBとモジュール構成 | `gorm`, Go Modules, DI設計など（任意） |

> 💡 各週末にミニプロジェクトを構築し、GitHubにコミットしていきます

---

## 🚀 進捗一覧

| Day | タイトル | 概要 | フォルダ |
|-----|----------|------|----------|
| 1 | Hello CLI | コマンドライン引数で挨拶するCLIツール | `/week1/hello-cli` |
| 2 | Text Analyzer | 入力された文章の文字数・単語数を解析するCLI | `/week1/text-analyzer` |
| ... | ... | ... | ... |

> 各プロジェクトは独立した `main.go` と README を持ち、単体でビルド＆実行できます

---

## 🛠️ 使用技術・ツール

- Go 1.20+
- `flag`, `net/http`, `encoding/json`
- `gorilla/mux` または `chi`
- `gorm`, SQLite（Week3以降）

---

## 🧠 比較視点：Laravelとの違い

| Laravel | Go |
|--------|----|
| Service/Repository パターン | interface + struct 構成 |
| Middleware | ハンドラチェイン＋明示的処理 |
| ORM（Eloquent） | gorm や ent（型安全） |
| Queue/Job | goroutine + channel |

> Laravel経験者が「Goらしさ」を掴めるように、設計視点も意識したトレーニングです

---


