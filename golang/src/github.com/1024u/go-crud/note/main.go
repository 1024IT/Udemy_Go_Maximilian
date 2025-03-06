package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {

	//ユーザからノートの値を取得
	title, content := getNoteData()

	//ユーザーからTODOの値を取得
	todoText := getUserInput("Todo text: ")

	//TODOのデータの入力チェック
	todo, err := todo.New(todoText)

	//TODOのデータの入力チェック
	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	//ノートのデータの入力チェック
	userNote, err := note.New(title, content)

	//ノートのデータの入力チェック
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

// interface{}...どんな種類のデータでも受け付けるデータ型
func printSomething(value interface{}) {

	//value.(int)...value が int 型かどうかを確認
	//ok...true なら value は int 型、false なら違う型
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	//value.(float64)...valueが float64型かどうかを確認
	//ok...true なら value は float64型、false なら違う型
	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float", floatVal)
		return
	}

	// value.name
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

// 「入力したデータの表示(Display())」と「セーブできたかどうかの判定(saveData)」を実施
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// セーブできたかどうかの判定メソッド
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

// 入力データを受けつけるメソッド
func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content:")

	return title, content
}

// func getTodoData() string {
// 	return getUserInput("Todo text: ")
// }

func getUserInput(prompt string) string {
	//%v..."fmt"パッケージにおけるフォーマット指定子の一種。構造体（複数の値を格納した変数的なもの）を出力する際、フィールド名が表示される(https://pkg.go.dev/fmt)
	fmt.Printf("%v ", prompt)

	//bufio.NewReader()...()内の値をバッファリングするリーダー
	//※バッファリング...一時的なメモリーにデータを蓄えておくこと（そこからデータを取り出したりできる）
	//os.Stdin()...ターミナルからの入力を受け付ける
	reader := bufio.NewReader(os.Stdin)

	//reader.ReadString...()内の値（区切り文字）が最初に現れるまでの文字列を返す
	// \n...フォーマット指定子の一種。改行する。
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	//strings.TrimSuffix...第一引数から、第二引数を取り除いた値を返す
	//\n...Linux/macOSにおける改行コード
	//\r...古いMac(〜)における改行コード
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
