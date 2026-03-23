package main

import (
	"bufio"
	"os"
	"strings"
)

type options struct {
	dislayReturn bool
	escapeEnable bool
}

func newOptions() *options {
	return &options{
		dislayReturn: true,
		escapeEnable: false,
	}
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	Run(os.Args[1:], w)
}

func Run(args []string, w *bufio.Writer) {
	options := newOptions()
	flags := validateArgs(args, w)
	switchOptions(flags, options)
	if len(flags) != 0 {
		args = append(args[:0], args[1:]...)
	}
	for i, val := range args {
		if i > 0 {
			w.WriteString(" ")
		}
		w.WriteString(val)
	}
	if options.dislayReturn {
		w.WriteString("\n")
	}
}

func validateArgs(args []string, w *bufio.Writer) []string {
	options := make([]string, 0, 3)
	if len(args) == 0 {
		w.WriteString("\n")
		return options
	}
	arg := args[0]
	starts := strings.HasPrefix(arg, "-")
	if starts {
		for _, v := range arg {
			if v == '-' {
				continue
			}
			switch v {
			case 'n', 'e', 'E':
				options = append(options, string(v))
			}
		}
	}
	return options
}

func switchOptions(flags []string, o *options) {
	for _, val := range flags {
		switch val {
		case "n":
			o.dislayReturn = false
		case "e":
			o.escapeEnable = true
		case "E":
			o.escapeEnable = false
		}
	}
}
