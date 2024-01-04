package main

import (
	"github.com/vultras/cli/mtool"
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
			"print string array to standard output",
		).Usage(
			"[str1 str2 ... strN]",
		),
		mtool.T("sum").Func(func(flags *mtool.Flags) {
			args := flags.Parse()
			one, _ := strconv.Atoi(args[0])
			two, _ := strconv.Atoi(args[1])
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
				"first sub tool",
			),
			mtool.T("second").Func(func(flags *mtool.Flags) {
				fmt.Println("called the second", flags.Parse())
			}).Desc(
				"second sub tool",
			),
		).Desc(
			"the tool to demonstrate how subtools work",
		),
	).Desc(
		"the testing program to show how to use the lib",
	).Ldesc(
		"this is the long description where you " +
			"can put anything you want about the program",
	)
)

func main() {
	root.Run(os.Args[1:])
}
