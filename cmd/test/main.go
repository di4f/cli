package main

import (
	"github.com/omnipunk/cli/mtool"
	"strconv"
	"fmt"
	"os"
)

var (
	root = mtool.T("test").Subs(
		mtool.T("echo").Func(func(flags *mtool.Flags) {
				var b bool
				flags.BoolVar(&b, "b", false, "the check flag")
				args := flags.Parse()
				fmt.Println(args)
			}).Desc(
				"print string to standard output string",
			).Usage(
				"[str1 str2 ... strN]",
			),
		mtool.T("sum").Func(func(flags *mtool.Flags) {
				flags.Parse()
				args := flags.Args()
				one, _ := strconv.Atoi(args[1])
				two, _ := strconv.Atoi(args[2])
				fmt.Println(one + two)
			}).Desc(
				"add one value to another",
			).Usage(
				"<int1> <int2>",
			),
		mtool.T("sub").Subs(
			mtool.T("first").Func(func(flags *mtool.Flags) {
					fmt.Println("called the first", flags.Parse())
				}).Desc(
					"description",
				).Usage(
					"[nothing here]",
				),
			mtool.T("second").Func(func(flags *mtool.Flags){
					fmt.Println("called the second", flags.Parse())
				}).Desc(
					"description",
				).Usage(
					"[nothing here]",
				),
			),
	)
)

func main() {
	root.Run(os.Args[1:])
}
