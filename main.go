// Author: Diego Emilio Moreno Sanchez
// This program is a simple implementation of a CPU scheduler
// using the Round Robin algorithm and the shortest job first
// algorithm.
package main

import "fmt"

func main() {

  // Options
  var option int

  for {
    menu()
    fmt.Scanln(&option)
    switch option {
    case 1:
      shortestJobFirst()
    case 2:
      roundRobin()
    case 3:
      return
    default:
      fmt.Println("Invalid option")
    }
  }
}

