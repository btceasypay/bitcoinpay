package bx

import (
	"fmt"
	"os"
)

func ErrExit(err error) {
	fmt.Fprintf(os.Stderr, "Bx Error : %q\n", err)
	os.Exit(1)
}
