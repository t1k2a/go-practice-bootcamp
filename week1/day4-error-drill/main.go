// 🎯 ゴール
// Goのエラーハンドリング（if err != nil）に慣れる

// Laravelでいう「例外クラス」「throw/catch」との違いを体験

// errors.New, fmt.Errorf, カスタムエラー型を使ってみる

package main

import (
	// "errors"
	"fmt"
	"strings"
)

func Greet(name string) (string, error) {
	errorStrings := []string{}
	errorStrings = append(errorStrings, nameValidation(name)...)


	if len(errorStrings) > 0 {
		return "", fmt.Errorf("バリデーションエラー - %s", strings.Join(errorStrings, ", "))
	}

	return fmt.Sprintf("こんにちは、%sさん！", name), nil
}

// チャレンジ課題（任意）　
// 複数のエラーを []error で返してみる（Goではよくあるパターン）
func nameValidation(name string) []string  {
	errorStrings := []string{}
	if name == "" {
		errorStrings = append(errorStrings, "名前が空です")
	}

	if strings.ContainsAny(name, "1234567890") {
		errorStrings = append(errorStrings, "名前に数字が含まれています")
	}

	return errorStrings
}

// カスタムエラー型（構造体ベース）
type ValidationError struct {
	Field string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("バリデーションエラー - %s: %s", e.Field, e.Message)
}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return &ValidationError {
			Field: "email",
			Message: "メールアドレスに@が含まれていません",
		}
	}

	return nil
}

func main() {
	// Greet関数テスト
	names := []string{"Taro", "", "Jiro123"}
	for _, name := range names {
		msg, err := Greet(name)
		if err != nil {
			fmt.Println("エラー：", err)
			continue
		}

		fmt.Println(msg)
	}

	// Emailバリデーションテスト
	emails := []string{"example@example.com", "invalid-email"}
	for _, email := range emails {
		err := ValidateEmail(email)
		if err != nil {
			// 型アサーションでカスタムエラーの詳細取得
			if ve, ok := err.(*ValidationError); ok {
				fmt.Printf("フィールド: %sのエラー → %s \n", ve.Field, ve.Message)
			} else {
				fmt.Println("一般エラー:", err)
			}
		} else {
			fmt.Printf("有効なメール: %s\n", email)
		}
	}
}