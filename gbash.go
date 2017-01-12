package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	//"path"
	//"runtime"
	"bytes"
	"regexp"
	"strings"
)

var testMode bool

func initTempFiles() (args, ins, outs, errs string, err error) {
	var argf, inf, outf, errf *os.File
	tmpdir := "c:\\Temp\\"
	arg := tmpdir + "_args"

	b, err := ioutil.ReadFile(arg)
	if err != nil {
		return
	}
	err = os.Remove(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
	arg = strings.TrimSpace(string(b))
	testMode = strings.HasPrefix(arg, "-t ")
	if testMode {
		arg = strings.TrimSpace(arg[3:])
		fmt.Println(arg)
	}
	if len(arg) > 2 {
		if (strings.Count(arg, "'") == 2 && arg[0] == '\'' && arg[len(arg)-1] == '\'') ||
			(strings.Count(arg, "\"") == 2 && arg[0] == '"' && arg[len(arg)-1] == '"') {
			arg = strings.TrimSpace(arg[1:len(arg)])
		}
	}
	argf, err = ioutil.TempFile("", "_args")
	if err != nil {
		return
	}
	inf, err = ioutil.TempFile("", "_stdin")
	if err != nil {
		return
	}
	outf, err = ioutil.TempFile("", "_stdout")
	if err != nil {
		return
	}
	errf, err = ioutil.TempFile("", "_stderr")
	if err != nil {
		return
	}

	args = argf.Name()
	ins = inf.Name()
	outs = outf.Name()
	errs = errf.Name()

	lins := " <" + linuxPath(ins)
	if i := strings.Index(arg, "<"); i > 0 && arg[i-1] != '\\' {
		lins = ""
	}
	louts := " >" + linuxPath(outs)
	if i := strings.Index(arg, ">"); i > 0 && arg[i-1] != '\\' {
		louts = ""
	}
	lerrs := " 2>" + linuxPath(errs)
	if strings.Contains(arg, "2>") {
		lerrs = ""
	}

	_, err = argf.WriteString(argsopt + linuxPath(arg) + lins + louts + lerrs)
	if err != nil {
		return
	}

	stat, e := os.Stdin.Stat()
	if e == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		// stdin has some data
		stdin := io.Reader(os.Stdin)
		b, err := ioutil.ReadAll(stdin)
		if err != nil {
			fmt.Println(err)
		}
		_, err = inf.Write(b)
	}

	argf.Close()
	inf.Close()
	outf.Close()
	errf.Close()

	return
}

func linuxPath(path string) (rst string) {
	re := regexp.MustCompile("(^|[='\"` ])([a-zA-Z]):\\\\")
	path = re.ReplaceAllStringFunc(path, func(s string) string {
		x := []byte{s[0]}
		if bytes.ContainsAny(x, "='\"` ") {
			return fmt.Sprintf("%s/mnt/%s/", x, bytes.ToLower([]byte{s[len(s)-3]}))
		} else {
			return fmt.Sprintf("/mnt/%s/", bytes.ToLower([]byte{s[len(s)-3]}))
		}
	})
	rst = strings.Replace(path, "\\", "/", -1)
	return
}
func main() {
	// get current path
	//_, currentFilePath, _, _ := runtime.Caller(0)
	//dirpath := path.Dir(currentFilePath)[3:]

	// make temp files
	argf, inf, outf, errf, err := initTempFiles()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	/*
		b, err := ioutil.ReadFile(inf.Name())
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			return
		}
		// drive root replacement
		for i, a := range os.Args[1:] {
			if strings.Index(a, ":\\") == 1 {
				os.Args[i+1] = "/mnt/" + strings.ToLower(string(a[0])) + "/" + a[3:]
				os.Args[i+1] = strings.Replace(os.Args[i+1], "\\", "/", -1)
			}
		}
		_, err = inf.WriteString(fmt.Sprintf(argsopt+"%s 2>%s", strings.Join(os.Args[1:], " "), errf.Name()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			return
		}*/

	execopt := []string{os.Getenv("SYSTEMROOT") + "\\system32\\gbash.vbs", linuxPath(argf)}
	err = exec.Command("isconemu").Run()
	if err == nil {
		execopt = append(execopt, "-cur_console:p")
	}

	cmd := exec.Command("wscript.exe", execopt...) //fmt.Sprintf("/mnt/c/%s/%s", dirpath, inf.Name())).Run()
	/*cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr*/

	/*in, err := cmd.StdinPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}*/

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		//return
	} else {
		err = cmd.Wait()
		if err != nil {
			fmt.Println(err)
			//fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			//return
		}
	}
	// print stdout
	b, err := ioutil.ReadFile(outf)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}
	if len(b) > 0 {
		fmt.Fprintf(os.Stdout, "%s\n", string(b))
	}

	// print stderr
	b, err = ioutil.ReadFile(errf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	if len(b) > 0 {
		fmt.Fprintf(os.Stderr, "%s\n", string(b))
	}

	// remove all temp files
	if !testMode {
		for _, f := range []string{argf, inf, outf, errf} {
			err := os.Remove(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			}
		}
	}

	return
	/*
		in, err := cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	/*out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}*/

	//stdin := io.Reader(os.Stdin)
	/*stdout := io.Reader(os.Stdout)
	go func(out io.Reader, in io.WriteCloser) {
		var b []byte
		for {
			n, e := out.Read(b)
			fmt.Println(n, e)
			if n > 0 && e == nil {
				in.Write(b)
			}
		}
	}(stdout, in)
	defer in.Close()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	cmd.Wait()*/
}
