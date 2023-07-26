package main

import (
	"database/sql"
	"fmt"
	"sync"
)

func main() {
	// defer文を有効活用しよう
	defer fmt.Println("world")
	fmt.Println("hello")

	// deferでsql.DB型のcloseをする
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// deferでロック解除をする
	var mu sync.Mutex
	src := []int{1, 2, 3}
	dst := []int{}

	for _, s := range src {
		go func(s int) {
			mu.Lock()
			defer mu.Unlock()

			dst = append(dst, s*2)
		}(s)
	}

	// deferでチャンネルをcloseする
	ch := make(chan int)
	go func(ch chan<- int) {
		defer close(ch)
	}(ch)

	fmt.Println(<-ch)

}

//go:generate stringer -type=Animal
type Animal int

const (
	Unknown Animal = iota
	Gopher
	Dog
	Cat
)

// enumに入る値のバリデーションを行う
func (a Animal) IsValid() bool {
	switch a {
	case Unknown, Gopher, Dog, Cat:
		return true
	}
	return false
}
