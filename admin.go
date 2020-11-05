package main

import (
    "fmt"
    "time"
    "./proceso"
)

var flag bool = false

func eliminaProceso(prcs *[]proceso.Proceso, s uint64) []proceso.Proceso {
    _, pos := isIn(prcs, s)
    return append(prcs[:pos], prcs[pos+1:])
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
    /*pro := proceso.Proceso{uint64(0), true}
    pro1 := proceso.Proceso{uint64(1), true}
    pros := []proceso.Proceso{}
    pros = append(pros, pro)
    pros = append(pros, pro1)
    c1 := make(chan uint64)

    for _, v := range pros {
        go v.Start(c1)
        go ProcesoF(v.Id)
        fmt.Println("FOR")
    }*/





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
                pro := proceso.Proceso{prc, true}
                c := make(chan uint64)
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
            var id uint64
            fmt.Print("Ingrese ID: ")
            fmt.Scan(&id)
            prcs = eliminaProceso(&prcs, id)
        }

        if opc == 4 {
            break
        }
    }


    var input string
    fmt.Scanln(&input)
}
