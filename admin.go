package main

import (
    "fmt"
    "time"
    "./proceso"
)

var flag bool = false

func eliminaProceso(prcs *[]proceso.Proceso, s int) {
    (*prcs)[s].Stop()
    *prcs = append((*prcs)[:s], (*prcs)[s+1:]...)
}

func isIn(prcs *[]proceso.Proceso, id uint64) (bool, int) {
    for i, v := range *prcs {
        if v.Id == id {
            return true, i
        }
    }
    return false, 0
}

func main() {
    var opc int
    prcs := []proceso.Proceso{}

    for {
        fmt.Println("1. Agregar Proceso")
        fmt.Println("2. Mostrar Procesos")
        fmt.Println("3. Terminar Proceso")
        fmt.Println("4. Salir")
        fmt.Scan(&opc)

        switch opc {
        case 1:
            var prc uint64
            fmt.Print("Ingrese ID: ")
            fmt.Scan(&prc)
            if cond, _ := isIn(&prcs, prc); cond {
                fmt.Println("\tId en uso\n")
            } else {
                c := make(chan bool)
                pro := proceso.Proceso{prc, 0, true, false, c}
                go pro.Start()
                prcs = append(prcs, pro)
            }
        case 2:
            var input string
            for _, p := range prcs {
                p.ActivoC <- true
            }
            fmt.Scanln(&input)
            for _, p := range prcs {
                p.ActivoC <- false
            }
        case 3:
            var id uint64
            fmt.Print("Ingrese ID: ")
            fmt.Scan(&id)
            _, v := isIn(&prcs, id)
            eliminaProceso(&prcs, v)
        }

        if opc == 4 {
            break
        }
    }
}
