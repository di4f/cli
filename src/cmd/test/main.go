package main

import(
	"github.com/mojosa-software/gomtool/src/mtool"
	"strconv"
	"fmt"
)

var(
	tools = mtool.Tools{
		"echo" : mtool.Tool{
			func(flags *mtool.Flags) {
				var b bool
				flags.BoolVar(&b, "b", false, "the check flag")
				flags.Parse()
				
				args := flags.Args()
				
				fmt.Println(args)
			},
			"print string to standard output string",
			"[str1 str2 ... strN]",
		},
		"sum" : mtool.Tool{
			func(flags *mtool.Flags) {
				flags.Parse()
				args := flags.Args()
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

