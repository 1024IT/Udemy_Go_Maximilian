package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// (note Note)...レシーバ引数。Note structの値を取得している。
func (note Note) Display() {
	//\n...Linux/macOSにおける改行コード
	//%v..."fmt"パッケージにおけるフォーマット指定子の一種。構造体（複数の値を格納した変数的なもの）を出力する際、フィールド名が表示される(https://pkg.go.dev/fmt)
	//		(%v %v, 値, 値)なので、
	fmt.Printf("Your note titled %v has the following content :\n\n%v\n\n", note.Title, note.Content)
}

// (note Note)...レシーバ引数。Note structの値を取得している。
func (note Note) Save() error {

	//strings.ReplaceAll()...指定した文字列内の特定の文字を変更する。
	// 　　　　　　　　　　　　　第一引数において変更を行いたい文字列を指定します。
	// 　　　　　　　　　　　　　第二引数では、変更をする文字の指定を行います。
	// 　　　　　　　　　　　　　第三引数では、変更後の文字列の指定を行います。
	fileName := strings.ReplaceAll(note.Title, " ", "_")

	//strings.ToLower...大文字を小文字に変換する。
	fileName = strings.ToLower(fileName) + ".json"

	//json.Marshal...構造体(Note structs)をjsonに変換する
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	//os.WriteFile...指定したファイルにデータの内容を書き込むメソッド
	//第一引数...データを書き込むファイル
	//第二引数...書き込むデータ
	//第三引数...ファイルの権限(0644...所有者が読み書きでき、グループと他のユーザーは読み取り専用)
	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
