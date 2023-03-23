package main

import(
	mtool "github.com/surdeus/gomtool/src/multitool"
	"strconv"
	"fmt"
)

var(
	tools = mtool.Tools{
		"echo" : mtool.Tool{
			func(args []string, flags *mtool.Flags) {
				var b bool
				flags.BoolVar(&b, "b", false, "the check flag")
				flags.Parse(args)
				fmt.Println(flags.Args())
			},
			"print string to standard output string",
			"[str1 str2 ... strN]",
		},
		"sum" : mtool.Tool{
			func(args []string, flags *mtool.Flags) {
				flags.Parse(args)
				one, _ := strconv.Atoi(args[1])
				two, _ := strconv.Atoi(args[2])
				fmt.Println(one + two)
			},
			"add one value to another",
			"<int1> <int2>",
		},
	}
)

func main() {
	mtool.Main("test", tools)
}

