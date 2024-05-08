package main

import (
  "fmt"
  "math/rand"
  "text/tabwriter"
  "os"
)

type Process struct {
  id int 
  burstTime int 
  arrivalTime int 
  priority int 
}

func sortProcessQueue(processQueue []Process) []Process { 
  for i := 0; i < len(processQueue); i++ {
    for j := i + 1; j < len(processQueue); j++ {
      if processQueue[i].burstTime > processQueue[j].burstTime {
        processQueue[i], processQueue[j] = processQueue[j], processQueue[i]
      }
    }
  }
  return processQueue
}

func generateRandomProcess(numberOfProcesses int) (processes []Process, totalTime float64) {
  maxArrivalTime := 0
  for i := 0; i < numberOfProcesses; i++ {
    burstTime := rand.Intn(10)
    arrivalTime := rand.Intn(10)
    processes = append(processes, Process{
      id: i,
      burstTime: burstTime,
      arrivalTime: arrivalTime,
      priority: rand.Intn(10),
    })
    if arrivalTime > maxArrivalTime {
      maxArrivalTime = arrivalTime 
    }
    totalTime += float64(burstTime)
  }
  return processes, totalTime + float64(maxArrivalTime)
}

func enterProcessManually(numberOfProcesses int) (processes []Process, totalTime float64) {
  maxArrivalTime := 0
  for i := 0; i < numberOfProcesses; i++ {
    var process Process
    fmt.Println("Ingrese el tiempo de duracion del proceso ", i + 1, ": ")
    fmt.Scanln(&process.burstTime)
    fmt.Println("Ingresa el tiempo de llegada del proceso ", i + 1, ": ")
    fmt.Scanln(&process.arrivalTime)
    fmt.Println("Ingresa la prioridad del proceso ", i + 1, ": ")
    fmt.Scanln(&process.priority)
    process.id = i
    processes = append(processes, process)
    totalTime += float64(process.burstTime)
    if process.arrivalTime > maxArrivalTime {
      maxArrivalTime = process.arrivalTime
    }
  }
  return processes, totalTime + float64(maxArrivalTime)
}

func printProcesses(processes []Process, totalTime float64) {
  fmt.Println("=====================================")
  fmt.Println("Los procesos son: ")
  w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
  fmt.Fprintln(w, "ID\tDuracion del proceso\tTiempo de llegada\tPrioridad")
  for _, process := range processes {
    fmt.Fprintf(w, "%d\t%d\t%d\t%d\n", process.id, process.burstTime, process.arrivalTime, process.priority)
  }
  fmt.Fprintln(w, "Tiempo total\t", totalTime)
  w.Flush()
}

