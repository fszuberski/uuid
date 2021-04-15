package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func main() {
	num := parseArgs()
	for i := 0; i < num; i++ {
		fmt.Println(uuid.New())
	}
}

func parseArgs() int {
	args := os.Args[1:]

	if len(args) == 0 {
		return 1
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("invalid argument: %s", args[0])
	}

	if num < 0 {
		return 1
	}

	return num
}