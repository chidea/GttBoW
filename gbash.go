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
	"time"
)

var testMode bool

func initTempFiles() (argnm, innm, outnm, errnm string, err error) {
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
	if testMode = strings.HasPrefix(arg, "-t "); testMode {
		arg = strings.TrimSpace(arg[3:])
		fmt.Println(arg)
	}
	if len(arg) > 2 {
		if (arg[0] == '\'' && arg[len(arg)-1] == '\'') ||
			(arg[0] == '"' && arg[len(arg)-1] == '"') {
			arg = strings.TrimSpace(arg[1 : len(arg)-1])
		}
	}
	arg = linuxPath(arg)
	if testMode {
		fmt.Println(arg)
	}

	/* std stream files ready */
	var argf, inf, outf, errf *os.File
	var linnm, loutnm, lerrnm string

	if argf, err = ioutil.TempFile("", "_args"); err != nil {
		return
	} else {
		argnm = argf.Name()
	}
	if i := strings.Index(arg, "<"); i > 0 && arg[i-1] != '\\' {
	} else {
		if inf, err = ioutil.TempFile("", "_stdin"); err != nil {
			return
		}
		innm = inf.Name()
		linnm = " <" + linuxPath(innm)
		if j := strings.Index(arg, ";"); j >= 0 {
			// pipe stdin into first command
			arg = arg[:j] + linnm + arg[j:]
		} else {
			arg += linnm
		}
	}

	if i := strings.Index(arg, ">"); i > 0 && arg[i-1] != '\\' {
	} else {
		if outf, err = ioutil.TempFile("", "_stdout"); err != nil {
			return
		}
		outnm = outf.Name()
		loutnm = " >>" + linuxPath(outnm)
		arg = strings.Replace(arg, ";", loutnm+";", -1)
	}

	if strings.Contains(arg, "2>") {
	} else {
		if errf, err = ioutil.TempFile("", "_stderr"); err != nil {
			return
		}
		errnm = errf.Name()
		lerrnm = " 2>>" + linuxPath(errnm)
		arg = strings.Replace(arg, ";", lerrnm+";", -1)
	}

	_, err = argf.WriteString(argsopt + arg + loutnm + lerrnm)
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
func cp(c *bool, ch chan bool, w, r *os.File) {
	for *c {
		//fmt.Printf("%s", string(b))
		//ioutil.ReadAll(r)
		c, _ := io.Copy(w, r)
		//b, _ := ioutil.ReadAll(r)
		//if len(b) == 0 {
		if c == 0 {
			time.Sleep(1 * time.Millisecond)
		} // else {
		//fmt.Fprintf(w, "%s", string(b))
		//}
	}
	//w.Close()
	r.Close()
	ch <- true
}
func main() {
	// get current path
	//_, currentFilePath, _, _ := runtime.Caller(0)
	//dirpath := path.Dir(currentFilePath)[3:]

	// make temp files
	argnm, innm, outnm, errnm, err := initTempFiles()
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

	execopt := []string{os.Getenv("SYSTEMROOT") + "\\system32\\gbash.vbs", linuxPath(argnm)}
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
	c := true
	//chi := make(chan bool)
	cho := make(chan bool)
	che := make(chan bool)
	//inf, _ := os.OpenFile(outnm, os.O_WRONLY, 0666)
	//go cp(&c, chi, inf, os.Stdin)
	outf, _ := os.OpenFile(outnm, os.O_RDONLY, 0666)
	go cp(&c, cho, os.Stdout, outf)
	errf, _ := os.OpenFile(errnm, os.O_RDONLY, 0666)
	go cp(&c, che, os.Stderr, errf)

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "%s\n", err)
	} else {
		err = cmd.Wait()
		if err != nil {
			fmt.Fprintln(os.Stderr, "%s\n", err)
		}
	}
	c = false
	//<-chi
	<-cho
	<-che
	// print stdout
	/*b, err := ioutil.ReadFile(outnm)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}
	if len(b) > 0 {
		fmt.Fprintf(os.Stdout, "%s\n", string(b))
	}

	// print stderr
	b, err = ioutil.ReadFile(errnm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	if len(b) > 0 {
		fmt.Fprintf(os.Stderr, "%s\n", string(b))
	}*/

	// remove all temp files
	if testMode {
		for _, f := range []string{argnm, innm, outnm, errnm} {
			fmt.Println(f)
			b, err := ioutil.ReadFile(f)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%s\n", err.Error())
				return
			}
			if len(b) > 0 {
				fmt.Fprintf(os.Stdout, "%s\n", string(b))
			}
		}
	}
	for _, f := range []string{argnm, innm, outnm, errnm} {
		if f == "" {
			continue
		}
		err := os.Remove(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
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
