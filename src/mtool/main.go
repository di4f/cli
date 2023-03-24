package mtool

import(
	"fmt"
	"os"
	path "path/filepath"
	"flag"
	"sort"
)

type Flags struct {
	*flag.FlagSet
	progName string
	utilName string
	args []string
	parsedArgs []string
}

type Handler func(*Flags)

type Tool struct {
	Handler Handler
	Desc, Usage string
}

type Tools map[string] Tool

func (flags *Flags) Parse() {
	flags.FlagSet.Parse(flags.args)
	flags.parsedArgs = flags.FlagSet.Args()
}

func (flags *Flags) Args() []string {
	return flags.parsedArgs
}

func (flags *Flags) ProgName() string {
	return flags.progName
}

func (flags *Flags) UtilName() string {
	return flags.utilName
}

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
			keys := make([]string, len(m))
			i := 0
			for k, _ := range m {
				keys[i] = k
				i++
			}
			sort.Strings(keys)
			
			for _, k := range keys {
				tool := m[k]
				fmt.Printf("%s: %s\n", k, tool.Desc)
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
		FlagSet : flagSet,
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
	
	flags.progName = name
	flags.utilName = args[0]
	flags.args = args[1:]
	
	util.Handler(flags)
}

