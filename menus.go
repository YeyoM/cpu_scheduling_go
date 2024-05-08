package main

import "fmt"

// Menus 
func menu() {
  fmt.Println("=====================================")
  fmt.Println("Simulador de algoritmos de planificacion de procesos")
  fmt.Println("1. Shortest Job First (Proceso mas corto primero)")
  fmt.Println("2. Round Robin")
  fmt.Println("3. Salir")
  fmt.Print("Ingrese una opcion: ")
}

func menuProcess() {
  fmt.Println("=====================================")
  fmt.Println("Procesos")
  fmt.Println("1. Generar procesos aleatorios")
  fmt.Println("2. Ingresar procesos manualmente")
  fmt.Println("3. Atras")
  fmt.Print("Ingrese una opcion: ")
}
