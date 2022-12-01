package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <day>")
		return
	}

	arg := os.Args[1]

	// current time in milliseconds
	start := time.Now().UnixNano() / int64(time.Millisecond)

	switch arg {
	case "01":
		Day01()
	case "02":
		Day02()
	default:
		fmt.Println("Day not implemented")
	}

	// current time in milliseconds
	end := time.Now().UnixNano() / int64(time.Millisecond)

	// print execution time
	fmt.Printf("Execution time: %d ms", end-start)

}
