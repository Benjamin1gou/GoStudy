package main

import (
	"errors"
	"fmt"
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

}
