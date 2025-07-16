// 🎯 ゴール
// Goの構造体と外部ファイル（JSON, YAML）のマッピングに慣れる
// Laravelでいう .env や config/*.php に相当する 設定ファイルの扱い方 を体験
// encoding/json、gopkg.in/yaml.v2 の使い方を学ぶ

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"gopkg.in/yaml.v2"
	"strings"
)

type Config struct {
	AppName string `json:"app_name" yaml:"app_name"`
	Port int `json:"port" yaml:"port"`
	Debug bool `json:"debug" yaml:"debug"`
}

func loadConfig(path string) (*Config, error) {
	isYaml := strings.HasSuffix(strings.ToLower(path), ".yaml") ||
		strings.HasSuffix(strings.ToLower(path), ".yml")

	file, err := os.Open(path)
	if err  != nil {
		return nil, fmt.Errorf("設定ファイルの読み込み失敗: %w", err)
	}
	defer file.Close()

	var cfg Config

	if isYaml {
		// YAML処理
		data, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("設定ファイルの読み込み失敗: %w", err)
		}
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return nil, fmt.Errorf("設定ファイルのパース失敗: %w", err)
		}
	} else {
		// JSON処理
		if err := json.NewDecoder(file).Decode(&cfg); err != nil {
			return nil, fmt.Errorf("設定ファイルのパース失敗: %w", err)
		}
	}

	return &cfg, nil
}

func printConfig(filename string, config *Config) {
        fmt.Printf("=== %s ===\n", filename)
        fmt.Printf("アプリ名: %s\n", config.AppName)
        fmt.Printf("ポート: %d\n", config.Port)
        fmt.Printf("デバッグモード: %v\n", config.Debug)
        fmt.Println()
}

func loadAndPrint(filename string) {
	config, err := loadConfig(filename)
	if err != nil {
		fmt.Printf("エラー (%s): %v\n", filename, err)
		return
	}
	printConfig(filename, config)
}

// チャレンジ課題（任意）
//  YAML対応
//  go get gopkg.in/yaml.v2 して config.yaml も読めるようにする
func main() {
	files := []string{"config.json", "config.yaml"}
	for _, file := range files {
		loadAndPrint(file)
	}
}
