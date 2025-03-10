//modファイル...このプログラムで使うモジュール（部品）が記載されたファイル。
//             外部から取得したメソッドの取得元のパス（URL）が記載されている。

module example.com/bank

go 1.22.3

//go get {url}をターミナルで実行すれば、自動的に↓以下が記載される。
require github.com/Pallinder/go-randomdata v1.2.0 // indirect
