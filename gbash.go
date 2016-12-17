package main

import (
	"fmt"
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

func initTempFiles() (inf *os.File, errf *os.File, err error) {
	tmpin := "c:\\Temp\\_tmpin"

	b, err := ioutil.ReadFile(tmpin)
	if err != nil {
		return
	}
	err = os.Remove(tmpin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
	tmpin = strings.TrimSpace(string(b))
	testMode = strings.HasPrefix(tmpin, "-t ")
	if testMode {
		tmpin = strings.TrimSpace(tmpin[3:])
		fmt.Println(tmpin)
	}
	if len(tmpin) > 2 {
		if (strings.Count(tmpin, "'") == 2 && tmpin[0] == '\'' && tmpin[len(tmpin)-1] == '\'') ||
			(strings.Count(tmpin, "\"") == 2 && tmpin[0] == '"' && tmpin[len(tmpin)-1] == '"') {
			tmpin = strings.TrimSpace(tmpin[1:len(tmpin)])
		}
	}
	inf, err = ioutil.TempFile(".", "_stdin")
	if err != nil {
		return
	}
	errf, err = ioutil.TempFile(".", "_stderr")
	if err != nil {
		return
	}

	/*for {
		if idx := strings.Index(tmpin, ":\\"); idx > 0 {
			tmpin = tmpin[:idx-1] + "/mnt/" + strings.ToLower(string(tmpin[idx-1])) + "/" + tmpin[idx+2:]
		} else {
			break
		}
	}*/

	//tmpin = re.ReplaceAllString(tmpin, "${1}/mnt/${2}\\/")

	re := regexp.MustCompile("['\"` ][a-zA-Z]:\\\\")
	tmpin = re.ReplaceAllStringFunc(tmpin, func(s string) string {
		return fmt.Sprintf("%s/mnt/%s/", []byte{s[0]}, bytes.ToLower([]byte{s[1]}))
	})

	tmpin = strings.Replace(tmpin, "\\", "/", -1)

	_, err = inf.WriteString(stdinopt + tmpin + " 2>" + errf.Name())
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
	inf, errf, err := initTempFiles()
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
		_, err = inf.WriteString(fmt.Sprintf(stdinopt+"%s 2>%s", strings.Join(os.Args[1:], " "), errf.Name()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			return
		}*/
	inf.Close()

	errf.Close()

	execopt := []string{inf.Name()}
	err = exec.Command("isconemu").Run()
	if err == nil {
		execopt = append(execopt, "-cur_console:p")
	}

	cmd := exec.Command("bash", execopt...) //fmt.Sprintf("/mnt/c/%s/%s", dirpath, inf.Name())).Run()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

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

	b, err := ioutil.ReadFile(errf.Name())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	if len(b) > 0 {
		fmt.Fprintf(os.Stderr, "%s\n", string(b))
	}

	// remove all temp files
	if !testMode {
		for _, f := range []*os.File{inf, errf} {
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
