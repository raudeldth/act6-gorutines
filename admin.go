package main

import (
    "fmt"
    "time"
)

func Proceso(id uint64) {
    i := uint64(0);
    for {
        fmt.Printf("id %d: %d\t", id, i)
        i = i + 1
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    var opc int
    fmt.Println("1. Agregar Proceso")
    fmt.Println("2. Mostrar Procesos")
    fmt.Println("3. Terminar Proceso")
    fmt.Println("4. Salir")
    fmt.Scanf(&ocp)

    go Proceso(11)

    var input string
    fmt.Scanln(&input)
}
