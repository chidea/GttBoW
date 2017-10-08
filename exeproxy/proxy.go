package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmd string

func main() {
	os.Args[0] = cmd
	c := exec.Command("gbash.exe", os.Args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if e := c.Run(); e != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%e", e))
		os.Exit(1)
	}
}
