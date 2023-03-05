package main

import (
	"os"

	"github.com/ctopher7/gltc/v2/part1"
)

func main() {
	switch os.Args[1] {
	case "grpc":
		part1.ServeGrpc()
	case "consumer":
		part1.ServeConsumer()
	}
}
