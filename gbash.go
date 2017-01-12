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

func initTempFiles() (argf *os.File, inf *os.File, outf *os.File, errf *os.File, err error) {
	args := "c:\\Temp\\_args"

	b, err := ioutil.ReadFile(args)
	if err != nil {
		return
	}
	err = os.Remove(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
	args = strings.TrimSpace(string(b))
	testMode = strings.HasPrefix(args, "-t ")
	if testMode {
		args = strings.TrimSpace(args[3:])
		fmt.Println(args)
	}
	if len(args) > 2 {
		if (strings.Count(args, "'") == 2 && args[0] == '\'' && args[len(args)-1] == '\'') ||
			(strings.Count(args, "\"") == 2 && args[0] == '"' && args[len(args)-1] == '"') {
			args = strings.TrimSpace(args[1:len(args)])
		}
	}
	argf, err = ioutil.TempFile(".", "_args")
	if err != nil {
		return
	}
	inf, err = ioutil.TempFile(".", "_stdin")
	if err != nil {
		return
	}
	outf, err = ioutil.TempFile(".", "_stdout")
	if err != nil {
		return
	}
	errf, err = ioutil.TempFile(".", "_stderr")
	if err != nil {
		return
	}

	/*for {
		if idx := strings.Index(args, ":\\"); idx > 0 {
			args = args[:idx-1] + "/mnt/" + strings.ToLower(string(args[idx-1])) + "/" + args[idx+2:]
		} else {
			break
		}
	}*/

	//args = re.ReplaceAllString(args, "${1}/mnt/${2}\\/")

	re := regexp.MustCompile("[='\"` ][a-zA-Z]:\\\\")
	args = re.ReplaceAllStringFunc(args, func(s string) string {
		return fmt.Sprintf("%s/mnt/%s/", []byte{s[0]}, bytes.ToLower([]byte{s[1]}))
	})

	args = strings.Replace(args, "\\", "/", -1)

	_, err = argf.WriteString(argsopt + args + " <" + inf.Name() + " >" + outf.Name() + " 2>" + errf.Name())
	if err != nil {
		return
	}

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

	execopt := []string{os.Getenv("SYSTEMROOT") + "\\system32\\gbash.vbs", argf.Name()}
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
	b, err := ioutil.ReadFile(outf.Name())
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}
	if len(b) > 0 {
		fmt.Fprintf(os.Stdout, "%s\n", string(b))
	}

	// print stderr
	b, err = ioutil.ReadFile(errf.Name())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	if len(b) > 0 {
		fmt.Fprintf(os.Stderr, "%s\n", string(b))
	}

	// remove all temp files
	if !testMode {
		for _, f := range []*os.File{argf, inf, outf, errf} {
			err := os.Remove(f.Name())
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
