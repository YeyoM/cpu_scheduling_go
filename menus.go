package main

import "fmt"

// Menus 
func menu() {
  fmt.Println("CPU Scheduler")
  fmt.Println("1. Shortest Job First")
  fmt.Println("2. Round Robin")
  fmt.Println("3. Exit")
  fmt.Print("Enter an option: ")
}

func menuProcess() {
  fmt.Println("Process")
  fmt.Println("1. Generate random processes")
  fmt.Println("2. Enter processes manually")
  fmt.Println("3. Back")
  fmt.Print("Enter an option: ")
}
