package main

import "fmt"
import "os"


// 入口 package is main
//     function is main

// main not support return value
// os.Exit() return status

// use os.Args to get parameters

func main() {
  fmt.Println(os.Args)
  if len(os.Args) > 1 {
    fmt.Println("hello world", os.Args[1])
  }
  os.Exit(255)
}
