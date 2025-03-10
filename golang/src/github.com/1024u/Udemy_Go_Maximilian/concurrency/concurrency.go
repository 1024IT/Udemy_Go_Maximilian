//学べること...
//Goroutines...Goにおける処理の最小単位
// 			   実現方法は、関数の前に「go」を入れることで、「go」と入れた関数を並行実行させる

package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)

	//チャネルから値trueを送信
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)

	//チャネルから値trueを送信
	doneChan <- true
}

func main() {

	dones := make([]chan bool, 4)
	//新しいチャネルを作成
	//done := make(chan bool)

	dones[0] = make(chan bool)
	go greet("Nice to meet you!", dones[0])
	dones[1] = make(chan bool)
	go greet("How are you?", dones[1])
	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", dones[2])
	dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", dones[3])

	//チャネルから値を受信
	for _, done := range dones {
		<-done
	}
}
