package main

import (
	"fmt"
	"os"
	"strings"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

func main() {
	argsWithoutProg := os.Args[1:]
	argsText := strings.Join(argsWithoutProg, " ")
	say, err := cowsay.Say(
		argsText,
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}
