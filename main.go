package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  status, msg := cli(os.Args[1:])
  fmt.Println(msg)
  os.Exit(status)
}

func cli(args []string) (status int, msg string) {
  status = 0
  if len(args) == 3 {
    if args[0] == "create" {
      numberOfInstances, err := strconv.Atoi(args[1])
      if err != nil || numberOfInstances <= 0 {
        return 1, "Error: Invalid number of instances"
      } else if numberOfInstances > 0 {
        if args[2] == "warrior" {
          msg = "Success: Warrior created"
        } else if args[2] == "archer" {
          msg = "Success: Archer created"
        } else {
          return 1, "Error: This plan is not supported"
        }
      }
    } else {
      return 1, "Error: Use create statement"
    }
  } else {
    return 1, "Error: 3 arguments needed"
  }
  
  return
}
