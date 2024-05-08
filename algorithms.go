package main

import (
  "fmt"
  "text/tabwriter"
  "os"
)

type ClockCycle struct {
  process Process
  time float64 
}

func shortestJobFirst() {
  var option int 
  var numberOfProcesses int
  for {
    menuProcess()
    fmt.Scanln(&option)
    switch option {
    case 1:
      fmt.Println("=====================================")
      fmt.Println("Proceso mas corto primero")
      fmt.Println("Ingrese el numero de procesos: ")
      fmt.Scanln(&numberOfProcesses)
      processes, totalTime := generateRandomProcess(numberOfProcesses)
      printProcesses(processes, totalTime)
      shortestJobFirstAlgorithm(processes, totalTime)
    case 2:
      fmt.Println("=====================================")
      fmt.Println("Proceso mas corto primero")
      fmt.Println("Ingrese el numero de procesos: ")
      fmt.Scanln(&numberOfProcesses)
      processes, totalTime := enterProcessManually(numberOfProcesses)
      printProcesses(processes, totalTime)
      shortestJobFirstAlgorithm(processes, totalTime)
    case 3:
      return
    default:
      fmt.Println("Opccion invalida")
    }
  }
}

func roundRobin() {
  var option int 
  var numberOfProcesses int
  var timeSlice int
  for {
    menuProcess()
    fmt.Scanln(&option)
    switch option {
    case 1:
      fmt.Println("=====================================")
      fmt.Println("Round Robin")
      fmt.Println("Ingrese el numero de procesos: ")
      fmt.Scanln(&numberOfProcesses)
      processes, totalTime := generateRandomProcess(numberOfProcesses)
      fmt.Println("Ingrese el tiempo de rafaga: ")
      fmt.Scanln(&timeSlice)
      printProcesses(processes, totalTime)
      roundRobinAlgorithm(processes, totalTime, timeSlice)
    case 2:
      fmt.Println("=====================================")
      fmt.Println("Round Robin")
      fmt.Println("Ingrese el numero de procesos: ")
      fmt.Scanln(&numberOfProcesses)
      processes, totalTime := enterProcessManually(numberOfProcesses)
      fmt.Println("Ingrese el tiempo de rafaga: ")
      fmt.Scanln(&timeSlice)
      printProcesses(processes, totalTime)
      roundRobinAlgorithm(processes, totalTime, timeSlice)
    case 3:
      return
    default:
      fmt.Println("Opccion invalida")
    }
  }
}

// Algorithms 
func shortestJobFirstAlgorithm(processes []Process, totalTime float64) {

  var currentCPUTime float64 = 0
  var clockCycle []ClockCycle 
  var processQueue []Process
  var isCPUIdle bool = true
  var nextFreeTime float64 = 0

  for currentCPUTime < totalTime {
    // Add the processes to the process queue 
    for _, process := range processes {
      if float64(process.arrivalTime) == currentCPUTime {
        processQueue = append(processQueue, process)
      }
    }

    if currentCPUTime >= nextFreeTime {
      isCPUIdle = true
    }

    // If the process queue is not empty
    if len(processQueue) > 0 && isCPUIdle {
      // Sort the process queue by burst time
      // and add the process to the clock cycle 
      // array 
      processQueue = sortProcessQueue(processQueue)
      clockCycle = append(clockCycle, ClockCycle{ 
        process: processQueue[0],
        time: currentCPUTime,
      })
      processQueue = processQueue[1:]
      isCPUIdle = false
      nextFreeTime = currentCPUTime + float64(clockCycle[len(clockCycle) - 1].process.burstTime)
    }

    currentCPUTime++
  }

  // Print the clock cycle array
  fmt.Println("Tabla de procesos")
  w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
  fmt.Fprintln(w, "ID\tTiempo de rafaga\tTiempo de llegada\tTiempo de inicio\tPrioridad")
  for _, clock := range clockCycle {
    fmt.Fprintf(w, "%d\t%d\t%d\t%.2f\t%d\n", clock.process.id, clock.process.burstTime, clock.process.arrivalTime, clock.time, clock.process.priority)
  }
  w.Flush()

}

/*
The round robin algorithm is a pre-emptive scheduling algorithm that assigns a fixed time unit per process. 
If the process is not finished within the time slice, the process is moved to the end of the queue. 
The algorithm is simple and easy to implement, but it is not the most efficient algorithm. 

The algorithm works as follows: 
1. Add the processes to the process queue. 
2. If the process queue is not empty, assign the process to the CPU. 
3. If the process is not finished within the time slice, move the process to the end of the queue. 
4. Repeat the process until all processes are finished. 
*/
func roundRobinAlgorithm(processes []Process, totalTime float64, timeSlice int) {
  
  var currentCPUTime float64 = 0
  var clockCycle []ClockCycle 
  var processQueue []Process
  var isCPUIdle bool = true
  var nextFreeTime float64 = 0

  for currentCPUTime < totalTime {
    // Add the processes to the process queue 
    for _, process := range processes {
      if float64(process.arrivalTime) == currentCPUTime {
        processQueue = append(processQueue, process)
      }
    }

    if currentCPUTime >= nextFreeTime {
      isCPUIdle = true
    }

    // If the process queue is not empty
    if len(processQueue) > 0 && isCPUIdle {
      // Add the process to the clock cycle array 
      clockCycle = append(clockCycle, ClockCycle{ 
        process: processQueue[0],
        time: currentCPUTime,
      })
      // If the process is not finished within the time slice
      if processQueue[0].burstTime > timeSlice {
        processQueue[0].burstTime -= timeSlice
        processQueue = append(processQueue, processQueue[0])
      }
      processQueue = processQueue[1:]
      isCPUIdle = false
      nextFreeTime = currentCPUTime + float64(timeSlice)
    }

    currentCPUTime++
  }

  // Print the clock cycle array
  fmt.Println("Tabla de procesos")
  w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
  fmt.Fprintln(w, "ID\tTiempo de rafaga\tTiempo de llegada\tTiempo de inicio\tPrioridad")
  for _, clock := range clockCycle {
    fmt.Fprintf(w, "%d\t%d\t%d\t%.2f\t%d\n", clock.process.id, clock.process.burstTime, clock.process.arrivalTime, clock.time, clock.process.priority)
  }
  w.Flush()

}
