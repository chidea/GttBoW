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
  fmt.Println("took", os.Args)
  for i, a := range os.Args{
    if i == 0 { continue }
    if strings.Contains(a, " ") {
      os.Args[i] = strings.Replace(a, " ", "\\ ", -1)
    }
    if strings.Contains(a, "\\") {
      if a[1:3] == ":\\" {
        a = "/mnt/" + strings.ToLower(string(a[0])) + "/" + a[3:]
      }
      os.Args[i] = strings.Replace(a, "\\", "/", -1)
    }
  }
  fmt.Println("call", os.Args)
  p := exec.Command("bash.exe", "-c", "git "+ strings.Join(os.Args[1:], " "))
  p.Stdin = os.Stdin
  p.Stdout = os.Stdout
  p.Stderr = os.Stderr
  e := p.Run()
  if e != nil {
    fmt.Println(e)
  }
}
