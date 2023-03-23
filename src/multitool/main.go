package multitool

import(
	"fmt"
	"os"
	path "path/filepath"
	"flag"
)

type Flags struct {
	*flag.FlagSet
}

type Handler func(args []string, flags *Flags)

type Tool struct {
	Handler Handler
	Desc, Usage string
}

type Tools map[string] Tool

func Main(name string, m Tools) {
	var(
		utilName string
		args []string
	)
	
	arg0 := os.Args[0]
	binBase := path.Base(arg0) ;
	binBase = binBase[:len(binBase)-len(path.Ext(binBase))]
	if binBase != name {
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
		fmt.Printf("%s: No such uitl as '%s'.\n", arg0, utilName )
		os.Exit(1)
	}

	util := m[utilName]
	
	arg1 := os.Args[1]
	flagSet := flag.NewFlagSet(arg1, flag.ExitOnError)
	flags := &Flags{
		flagSet,
	}
	flags.Usage = func() {
		out := flags.Output()
		n := 0
		flags.VisitAll(func(f *flag.Flag){
			n++
		})
		
		hasOptions := n != 0
		
		fmt.Fprintf(
			out,
			"Usage of %s:\n\t%s",
			arg1, arg1,
		)
		if hasOptions {
			fmt.Fprintf(out, " [options]")
		}
		
		if util.Usage != "" {
			fmt.Fprintf(
				out,
				" %s",
				util.Usage,
			)
		}
		
		fmt.Fprintln(out, "")
		
		if hasOptions {
			fmt.Fprintln(out, "Options:")
			flags.PrintDefaults()
		}
		
		os.Exit(1)
	}
	
	args = args[1:]
	
	util.Handler(args, flags)
}

