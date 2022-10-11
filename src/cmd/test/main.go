package main

import(
	"github.com/surdeus/gomtool/src/multitool"
	"strconv"
	"fmt"
)

var(
	tools = multitool.Tools{
		"echo" : multitool.Tool{
			func(args []string) {
				fmt.Println(args)
			},
			"print string to standard output string",
		},
		"sum" : multitool.Tool{
			func(args []string) {
				one, _ := strconv.Atoi(args[1])
				two, _ := strconv.Atoi(args[2])
				fmt.Println(one + two)
			},
			"add one value to another",
		},
	}
)

func main() {
	multitool.Main("test", tools)
}

