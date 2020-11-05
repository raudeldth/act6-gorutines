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
func ProcesoF(id uint64) {
    i := uint64(0);
    for {
        if flag {
            fmt.Printf("id %d: %d\n", id, i)
        }
        i = i + 1
        time.Sleep(time.Millisecond * 500)
    }
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
                c := make(chan uint64)
                pro := proceso.Proceso{prc, true}
                go pro.Start(c)
                go ProcesoF(pro.Id)
                prcs = append(prcs, pro)
            }
        case 2:
            flag = true
            var input string
            fmt.Scanln(&input)
            flag = false
        case 3:
            //FIXME
            var id uint64
            fmt.Print("Ingrese ID: ")
            fmt.Scan(&id)
            _, v := isIn(&prcs, id)
            eliminaProceso(&prcs, v)
            //FIXME
        }

        if opc == 4 {
            break
        }
    }

    var input string
    fmt.Scanln(&input)
}
