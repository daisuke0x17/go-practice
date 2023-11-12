package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum(c chan<- int) {
	fmt.Println("...")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	c <- num
}

func main() {
	fmt.Println("what is today's lucky number?")
	/*
		待ち合わせなしの場合
		go getLuckyNum()

		time.Sleep(time.Second * 5)
	*/

	// ============================================

	/*
		sync.WaitGroupで待機する場合
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			getLuckyNum()
		}()

		wg.Wait()
	*/

	// ============================================

	c := make(chan int)
	go getLuckyNum(c)

	num := <-c
	fmt.Printf("Today's your lucky number is %d", num)

	// 使い終わったチャネルはClose
	close(c)
}
