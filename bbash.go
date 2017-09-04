package main

import (
	//"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var debug bool

func main() {
	preargs := []string{}
	args := os.Args[1:]
	debug = len(args) > 0 && (args[0] == "-d" || args[0] == "--debug")
	if debug {
		debug = true
		args = args[1:]
		log.Println("argument before edit:", args)
	}
	err := exec.Command("isconemu").Run()
	if err == nil {
		preargs = append(preargs, "-cur_console:p")
	}
	var cmd *exec.Cmd
	if len(args) > 0 {
		for i := 0; i < len(args); i++ {
			args[i] = linuxPath(args[i])
		}
		args = []string{"-c", strings.Join(args, " ")}
		args = append(preargs, args...)
		cmd = exec.Command("bash.exe", args...)
	} else {
		cmd = exec.Command("bash.exe")
	}
	if debug {
		log.Println("arguments after edit:", args)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
}

var drivere = regexp.MustCompile(`(^|[='"` + "`" + ` ])(/[a-zA-Z]/|[a-zA-Z]:[/\\])`)

func linuxPath(path string) (rst string) {
	path = drivere.ReplaceAllStringFunc(path, func(s string) string {
		x := s[0]
		var y string
		if s[1] == ':' {
			y = string(s[0])
		} else {
			y = string(s[1])
		}
		y = strings.ToLower(y) // string(s[len(s)-3])))
		if debug {
			log.Println(s, x, y)
		}
		if strings.ContainsRune("='\"` ", rune(x)) {
			return fmt.Sprintf("%c/mnt/%s/", x, y)
		} else {
			return fmt.Sprintf("/mnt/%s/", y)
		}
	})
	rst = strings.Replace(path, "\\", "/", -1)
	return
}
