package main

import (
	"fmt"
	"os"
)

func main() {
	w := "world"
	// Sprint はフォーマットした結果を文字列でリターン
	hello := fmt.Sprintf("hello %s", w)
	fmt.Println(hello)

	// Fprint は書き込み先を明示的に指定
	f, err := os.Create("foo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(f, "Hello %s", w)
}
