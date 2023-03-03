package main

import (
	"go-lambda-hello/cooking"
	"go-lambda-hello/greeting"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func executeFunction() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) % 3
	switch num {
	case 0:
		greeting.SayHello()
	case 1:
		greeting.SayGoodMorning()
	case 2:
		cooking.SaySalmon()
	}
}

func main() {
	lambda.Start(executeFunction)
}
