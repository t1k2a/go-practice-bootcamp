// 🎯 ゴール
// CLIでタスクを登録・表示できるツールを作成
// JSONファイルにタスクを保存（永続化）
// --add や --list などのフラグで動作を切り替える

package main 

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const taskFile = "tasks.json"

type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	IsDone bool `json:"isDone"`
}

func getNextId(tasks []Task) int {
	if len(tasks) == 0 {
		return 1 // 明示的に1から開始
	}
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	
	return maxID + 1
}

func loadTasks() ([]Task, error) {
	file, err := os.Open(taskFile)
	if os.IsNotExist(err) {
		return []Task{}, nil // 初回起動時は空リスト
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(taskFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}

//  チャレンジ課題
// 🔹 タスクを完了済みにできるようにする（--done=2 など）
func main() {
	add := flag.String("add", "", "追加するタスク")
	list := flag.Bool("list", false, "タスク一覧表示")
	done := flag.Int("done", -1, "完了にするタスクのID")
	flag.Parse()

	if (*add != "" && *done != 1) || (*list && (*add != "" || *done != -1)) {
		fmt.Println("複数の操作は同時に指定できません。--add, --done, --list のうち1つだけ指定してください。")
		return
	}


	tasks, err := loadTasks()
	if err != nil  {
		fmt.Println("タスク読み込みエラー:", err)
		return
	}

	switch {
	case *add != "":
		tasks = append(tasks, Task{ID: getNextId(tasks), Description: *add, IsDone: false})
		if err := saveTasks(tasks); err != nil {
			fmt.Println("保存エラー:", err)
			return
		}

		fmt.Println("タスクを追加しました:", *add)

	case *list:
		if len(tasks) == 0 {
			fmt.Println("タスクはまだありません")
			return
		}
		fmt.Println("現在のタスク一覧:")
		for _, task := range tasks {
			fmt.Printf("%d. %s %t\n", task.ID, task.Description, task.IsDone)
		}
	case *done != -1:
		found := false
		for i, task := range tasks {
			if task.ID == *done {
				tasks[i].IsDone = true
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("ID %dのタスクが見つかりません\n", *done)
			return
		}
		if err := saveTasks(tasks); err != nil {
			fmt.Println("保存エラー:", err)
			return
		}
		fmt.Printf("ID %dのタスクを完了にしました\n", *done)
		

	default:
		fmt.Println("フラグを指定してください: --add, --list または --done")
	}
}
