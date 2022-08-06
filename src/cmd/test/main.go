package main

import(
	"github.com/surdeus/gomtool/src/multitool"
	"strconv"
	"fmt"
)

var(
	tools = multitool.Tools{
		"echo" : func(args []string) {
			fmt.Println(args)
		},
		"sum" : func(args []string) {
			one, _ := strconv.Atoi(args[1])
			two, _ := strconv.Atoi(args[2])
			fmt.Println(one + two)
		},
	}
)

func main() {
	multitool.Main("test", tools)
}

