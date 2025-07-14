// 🎯 ゴール
// Goの struct と Laravelのクラス（Eloquentモデルなど）を 手を動かしながら比較 して理解する。

// Goで User モデル風の構造体を定義

// メソッド・初期化・ポインタレシーバなどの書き方を体験

// LaravelのEloquentと何が違うのかを自分の手で確認

package main

import (
	"fmt"
	"strings"
)

type User struct {
	ID int
	Name string
	Email string
}

// 値レシーバ：読み取り用
// 読み取り専用、小さな構造体
func (u User) Greet() string {
	return fmt.Sprintf("こんにちは、%s さん！", u.Name)
} 

// ポインタレシーバ
// 値を変更する、大きな構造体
func (u *User) NormalizeEmail() {
	u.Email = strings.ToLower(u.Email)
}

// チャレンジ課題（任意）
// 複数の User を slice で持ち、NormalizeAll() を実装
func NormalizeAll(users []*User) {
	for _,user := range users {
		user.Name = strings.ToLower(user.Name)
		user.Email = strings.ToLower(user.Email)

		fmt.Printf("正規化後: %s - %s\n", user.Name, user.Email)
	}
}

func main() {
	user1 := User{
		ID: 1,
		Name: "Taro",
		Email: "TARO@EXAMPLE.COM",
	}

	user2 := User{
		ID: 2,
		Name: "Jiro",
		Email: "JIRO@EXAMPLE.COM",
	}

	fmt.Println(user1.Greet())
	fmt.Println()
	user1.NormalizeEmail()

	
	fmt.Println("正規化されたメール：", user1.Email)

	users := []*User{&user1, &user2}
	NormalizeAll(users)
	fmt.Println("正規化完了")

}