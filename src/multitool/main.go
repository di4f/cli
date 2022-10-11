package multitool

import(
	"fmt"
	"os"
	"path"
)

type Tool struct {
	Handler func(args []string)
	Desc string
}


type Tools map[string] Tool

func Main(name string, m Tools) {
	var(
		utilName string
		args []string
	)

	if binBase := path.Base(os.Args[0]) ; binBase != name {
		utilName = binBase
		args = os.Args[:]
	} else {
		if len(os.Args)<2  {
			for k, v := range m {
				fmt.Printf("%s:\t%s\n", k, v.Desc)
			}
			os.Exit(0)
		}
		utilName = os.Args[1]
		args = os.Args[1:]
	}

	if _, ok := m[utilName] ; !ok {
		fmt.Printf("%s: No such uitl as '%s'.\n", os.Args[0], utilName )
		os.Exit(1)
	}

	m[utilName].Handler(args)
}

