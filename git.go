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
  p := exec.Command("bash.exe", "-c", "git "+ strings.Join(os.Args[1:], " "))
  p.Stdin = os.Stdin
  p.Stdout = os.Stdout
  p.Stderr = os.Stderr
  e := p.Run()
  if e != nil {
    fmt.Println(e)
  }
}
