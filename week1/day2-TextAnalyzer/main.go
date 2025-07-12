// 🎯 ゴール
// 標準入力やフラグで受け取ったテキストから、以下を解析して出力するCLIツールを作成します：

// 文字数（全角含む）

// 単語数（空白区切り）

// 行数（改行区切り）

// Laravelでいうと「バリデーションの前処理」や「ログ整形」みたいなユーティリティに近い処理です。

package main

import (
	"bufio" // Bufferd I/O バッファ（一時的なメモリ領域）を使って効率的に入出力を使う
			// 小さな読み書きを何度も繰り返す代わりに、まとめて処理
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// フラグ定義：--text="解析する文字列"
	input := flag.String("text", "", "解析する文字列")
	flag.Parse()

	var text string

	if *input != "" {
		text = *input
	} else {
		// 標準入力から読み込み
		fmt.Println("解析するテキストを入力してください（Ctrl+Dで終了）:")
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		// 標準入力空手キスをと読み込み
		for scanner.Scan() {
			// １行ずつ読み込み、変数に追加していく
			lines = append(lines, scanner.Text())

			// Ctrl+D（EOF）が押されるとループが終了
		}

		// linesの\nごとにスライスして最後に開業文字列で結合
		text = strings.Join(lines, "\n")
	}

	// 文字数をカウント
	// runeを使うことでint32型の文字数を返す
	// 例：s := "Hello世界"
	// fmt.Println(len(s))           // 11バイト
	// fmt.Println(len([]rune(s)))   // 7文字
	charCount := len([]rune(text)) // Unicode対応
	// 空白区切りの単語をカウント
	wordCount := len(strings.Fields(text))
	// 改行区切りの行数をカウント
	lineCount := len(strings.Split(text, "\n"))

	fmt.Println("解析結果:")
	fmt.Printf("文字数: %d\n", charCount)
	fmt.Printf("単語数: %d\n", wordCount)
	fmt.Printf("行数: %d\n", lineCount)
}