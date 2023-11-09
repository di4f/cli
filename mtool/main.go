package mtool

// The package implements fast way
// to make multitool CLI applications.

import (
	"fmt"
	"os"
	//path "path/filepath"
	"flag"
	"sort"
)


type Flags struct {
	*flag.FlagSet
	tool *Tool
	args []string
	parsedArgs []string
}

type Handler interface {
	Handle(*Flags)
}

type HandlerFunc func(*Flags)
func (fn HandlerFunc) Handle(flags *Flags) {
	fn(flags)
}

type Tool struct {
	name string
	handler Handler
	desc, usage string
	subs ToolMap
	parent *Tool
}

// Returns new empty tool with specified name.
func T(name string) *Tool {
	ret := &Tool{}
	ret.name = name
	return ret
}

func (t *Tool) Handler(h Handler) *Tool {
	t.handler = h
	return t
}

func (t *Tool) Func(fn HandlerFunc) *Tool {
	t.handler = fn
	return t
}

func (t *Tool) Desc(d string) *Tool {
	t.desc = d
	return t
}

func (t *Tool) Usage(u string) *Tool {
	t.usage = u
	return t
}

func (t *Tool) Subs(tools ...*Tool) *Tool {
	if t.subs == nil {
		t.subs = ToolMap{}
	}
	for _, tool := range tools {
		tool.parent = t
		t.subs[tool.name] = tool
	}
	return t
}

func (t *Tool) Name() string {
	return t.name
}

func (t *Tool) FullName() string {
	ret := ""
	for t != nil {
		ret = t.name + func() string {
			if ret != "" {
				return " "
			}
			return ""
		}() + ret
		t = t.parent
	}
	return ret
}

func (t *Tool) ProgName() string {
	for t.parent != nil {
		t = t.parent
	}
	return t.name
}

func (t *Tool) Run(args []string) {
	var(
		usageTool *Tool
	)
	// Should implement the shit later
	//binBase := path.Base(arg0) ;
	//binBase = binBase[:len(binBase)-len(path.Ext(binBase))]

	flagSet := flag.NewFlagSet(t.FullName(), flag.ExitOnError)
	flags := &Flags{
		FlagSet : flagSet,
	}
	out := flags.Output()
	flags.Usage = func() {
		n := 0
		// We can visit the flags since the
		// function will be called after 
		// parsing.
		flags.VisitAll(func(f *flag.Flag){
			n++
		})
		
		hasOptions := n != 0
		
		fmt.Fprintf(
			out, "Usage:\n  %s",
			usageTool.FullName(),
		)
		if hasOptions {
			fmt.Fprintf(out, " [options]")
		}
		
		if usageTool.usage != "" {
			fmt.Fprintf(
				out,
				" %s",
				usageTool.usage,
			)
		}

		fmt.Fprintln(out, "")
		
		if hasOptions {
			fmt.Fprintln(out, "\nOptions:")
			flags.PrintDefaults()
		}
		
		os.Exit(1)
	}
	
	flags.args = args

	// If the tool has its own handler run it.
	if t.handler != nil {
		usageTool = t
		t.handler.Handle(flags)
		return
	}

	// Print available sub commands if
	// got no arguments.
	if len(args) == 0 {
		keys := make([]string, len(t.subs))
		i := 0
		for k, _ := range t.subs {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		
		fmt.Fprintf(out, "Usage:\n"+
			"  %s <command> [options] [arguments]\n\n" +
			"Commands:\n", t.FullName())
		for i, k := range keys {
			tool := t.subs[k]
			fmt.Fprintf(out, "  %s\t%s\n", k, tool.desc)
			if i != len(keys) - 1 {
				fmt.Fprintln(out, "")
			}
		}
		
		os.Exit(1)
	}
	toolName := args[0]
	args = args[1:]
	
	if _, ok := t.subs[toolName] ; !ok {
		fmt.Printf("%s: No such util %q'\n", t.ProgName(), toolName)
		os.Exit(1)
	}

	sub := t.subs[toolName]
	usageTool = sub
	sub.Run(args)
}

type ToolMap map[string] *Tool

func (flags *Flags) Parse() []string {
	flags.FlagSet.Parse(flags.args)
	flags.parsedArgs = flags.FlagSet.Args()
	return flags.parsedArgs
}

func (flags *Flags) AllArgs() []string {
	return flags.args
}

func (flags *Flags) Args() []string {
	return flags.parsedArgs
}

func (flags *Flags) Tool() *Tool {
	return flags.tool
}

