package model

type Todo struct {
	// 💡 チャレンジ課題（任意）
	// title に最大文字数制限を追加：validate:"required,max=20"
	Title string `json:"title" validate:"required,max=20"`
	Done bool `json:"done"`
}