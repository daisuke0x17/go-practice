package main

import (
	"fmt"
	"time"
)

func f(v string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %s\n", i, v)
		time.Sleep(3 * time.Second)
	}
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// go f("goroutineで実行")
	// f("普通に実行")
	// fmt.Println("終わり")

	// fmt.Println("--------------------------------------")

	// messages := make(chan string)
	// go func() {
	// 	messages <- "Hello"
	// }()

	// msg := <-messages
	// fmt.Println(msg)

	// fmt.Println("--------------------------------------")

	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// // ch <- 3

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// fmt.Println("--------------------------------------")

	// pings := make(chan string, 1)
	// pongs := make(chan string, 1)
	// ping(pings, "Hello")
	// pong(pings, pongs)
	// fmt.Println(<-pongs)

	fmt.Println("--------------------------------------")

	start := time.Now()

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println("result", <-results)
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
