package main

import (
	"bufio"
	"os"
)

var dislayReturn = true

func main() {
	args := os.Args[1:]
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	if len(args) == 0 {
		w.WriteString("\n")
		return
	}
	if args[0] == "-n" {
		dislayReturn = false
		args = append(args[:0], args[1:]...)
	}
	for i, val := range args {
		if i > 0 {
			w.WriteString(" ")
		}
		w.WriteString(val)
	}
	if dislayReturn {
		w.WriteString("\n")
	}
}
