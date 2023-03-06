package main

import (
	"os"

	"github.com/ctopher7/gltc/v2/part1"
	"github.com/ctopher7/gltc/v2/part2"
)

func main() {
	switch os.Args[1] {
	case "grpc":
		part1.ServeGrpc(os.Args[2])
	case "consumer":
		part1.ServeConsumer(os.Args[2])
	case "part2":
		part2.Main()
	}
}
