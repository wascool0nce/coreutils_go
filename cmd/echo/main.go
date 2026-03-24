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
	if len(args) == 0 {
		w.WriteString("\n")
		return
	}
	options := newOptions()
	flags := validateFlags(args)
	if len(flags) != 0 {
		switchOptions(flags[0], options)
		args = args[1:]
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

func validateFlags(args []string) []string {
	options := make([]string, 0, 3)

	if strings.HasPrefix(args[0], "-") {
		for _, char := range args[0][1:] {
			switch char {
			case 'n':
				options = append(options, string(char))
				continue
			default:
				options = make([]string, 0, 3)
				return options
			}
		}
	}
	return options
}

func switchOptions(val string, o *options) {
	switch val {
	case "n":
		o.dislayReturn = false
	}
}
