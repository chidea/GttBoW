package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var debug bool

func main() {
	bashargs := []string{}
	oriargs := os.Args[1:]
	var args []string
	var preargs string
	debug = len(oriargs) > 0 && (oriargs[0] == "-d" || oriargs[0] == "--debug")
	if debug {
		debug = true
		oriargs = oriargs[1:]
		log.Println(len(oriargs), "arguments before edit:", oriargs)
	}
	err := exec.Command("isconemu").Run()
	if err == nil {
		bashargs = append(bashargs, "-cur_console:p")
	}
	b, err := exec.Command("tasklist", "/fi", "imagename eq ssh-agent").Output()
	if err == nil {
		if !bytes.Equal(b[:4], []byte("INFO")) {
			preargs = ". ~/.ssh/ssh-agent.sh;"
		}
	}
	var cmd *exec.Cmd
	if len(oriargs) > 0 {
		for i := 0; i < len(oriargs); i++ {
			oriargs[i] = strings.Replace(linuxPath(oriargs[i]), " ", "\\ ", -1)
		}
		args = append(bashargs, []string{"-c", preargs + strings.Join(oriargs, " ")}...)
		cmd = exec.Command("bash.exe", args...)
	} else {
		cmd = exec.Command("bash.exe")
	}
	if debug {
		log.Println(len(args), "arguments after edit:", args)
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
