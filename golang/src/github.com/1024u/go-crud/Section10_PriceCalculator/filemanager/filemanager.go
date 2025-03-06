package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	//os.Open...指定したファイルを開く
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	//defer...現在実行されている関数(ReadLines())の処理終了後に、
	// 		  指定した関数(Cliose())を実行させる制御構文
	defer file.Close()

	//bufio.NewScanner()...指定したファイルの中のテキストを1行ずつ読み込む
	scanner := bufio.NewScanner(file)

	var lines []string

	//scanner.Scan()...ファイルの次の行を読み込む
	//scanner.Text()...その行のテキストを取得
	//中身の値をスライスlinesに格納
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//スキャン中のエラーを検知したら変数errに格納
	err = scanner.Err()

	if err != nil {
		//file.Close()
		return nil, errors.New("Failed to read line in file.")
	}

	//file.Close()
	return lines, nil
}

// WriteJSON...
// 第一引数...保存するJSONファイルのパス
// 第二引数...prices.goのTaxIncludedPriceJob structsの全ての変数
// interface{}...どんな方でも受け付けられるデータ型
func (fm FileManager) WriteResult(data interface{}) error {
	//書き込み権限がついたファイルを生成
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	//json.NewEncoder...引数で指定された場所にjsonを出力するエンコーダーを作成
	//NewEncoder...インメモリ表現からバイト列表現への変換。
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	//エンコードに失敗したらエラー出力
	if err != nil {
		//file.Close()
		return errors.New("Failed to convert data to JSON.")
	}

	//file.Close()
	return nil
}

// inputPathとoutputPathを定義
func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
