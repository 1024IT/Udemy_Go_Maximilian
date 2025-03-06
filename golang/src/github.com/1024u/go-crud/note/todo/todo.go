package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

// (note Note)...レシーバ引数。Note structの値を取得している。
func (todo Todo) Display() {
	fmt.Println("\n", todo.Text, "\n")
}

// (note Note)...レシーバ引数。Note structの値を取得している。
func (todo Todo) Save() error {

	fileName := "todo.json"

	//json.Marshal...構造体(Note structs)をjsonに変換する
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	//os.WriteFile...指定したファイルにデータの内容を書き込むメソッド
	//第一引数...データを書き込むファイル
	//第二引数...書き込むデータ
	//第三引数...ファイルの権限(0644...所有者が読み書きでき、グループと他のユーザーは読み取り専用)
	return os.WriteFile(fileName, json, 0644)
}

// 入力エラーチェック
func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input.")
	}

	return Todo{
		Text: content,
	}, nil
}
