package main
import (
  "fmt"
  "os"
  "os/exec"
  _ "bytes"
  "strings"
  _ "syscall"
  _ "time"
)


func main () {
  fmt.Println("calling git with", os.Args)
  for i, a := range os.Args{
    if strings.Contains(a, " ") {
      os.Args[i] = strings.Replace(a, " ", "\\ ", -1)
    }
  }
  /*attr := &os.ProcAttr { Dir : "", Env: nil}
  p, e := os.StartProcess("powershell", []string{"-Command", "ls"}, attr)
  if e != nil {
    fmt.Println(e)
  }else{
    p.Wait()
  }*/
  p := exec.Command("bash.exe", "-c", "git "+ strings.Join(os.Args[1:], " "))
  p.Stdin = os.Stdin
  p.Stdout = os.Stdout
  e := p.Run()
  if e != nil {
    fmt.Println(e)
    return
  }
  //time.Sleep(1 * time.Second)
  //e = p.Wait()
  if e != nil {
    fmt.Println(e)
  }
}
