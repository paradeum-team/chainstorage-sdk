package main

import (
	"fmt"
	"os"
)

func out(m string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, m, a...)
}

func processError(doing string, err error, args ...interface{}) {
	if err != nil {
		if len(args) > 0 {
			doing = fmt.Sprintf(doing, args...)
		}

		out("Error %s: %s\n", doing, err)

		os.Exit(1)
	}
}
