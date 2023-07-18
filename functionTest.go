package main

import (
	"errors"
	"fmt"
	"sync"
)

// 渡した名前を使用して、コマンドラインにテキストを出力する
func greet(name string) {
	fmt.Println("Hello", name)
}

// 割り算とあまり算出を行い2個の計算結果を返却
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// エラーハンドリング
func errorHundring(x float64) (float64, error) {
	if x < 0 {
		// 0とerror型の値を返却する
		return 0, errors.New("undefined for negative numbers")
	} else {
		return 1, nil
	}
}

func sayHello(wg *sync.WaitGroup) {
	fmt.Println("Hello")
	wg.Done() // Decrease counter
}

func sayWorld(wg *sync.WaitGroup) {
	fmt.Println("World")
	wg.Done() // Decrease counter
}

// mainファンクション
func main() {
	// テキスト出力関数の呼び出し
	greet("Alice")

	// 計算関数呼び出して、変数に格納
	quotient, remainder := divmod(7, 3)

	// 計算結果を出力
	fmt.Println("quotient:", quotient)
	fmt.Println("remainder:", remainder)

	// エラーハンドリング関数呼び出し
	number, message := errorHundring(1)
	// 値の存在チェック
	if message != nil {
		fmt.Println("Error:", message)
	} else {
		fmt.Println("Result:", number)
	}

	// ゴルーチン動作テスト
	var wg sync.WaitGroup

	wg.Add(2) // Increase counter

	go sayHello(&wg)
	go sayWorld(&wg)

	wg.Wait() // Wait for counter to be 0

}
