package proceso

import (
    "fmt"
    "time"
)

type Proceso struct {
    Id uint64
    I uint64
    Activo bool
    Write bool
    ActivoC chan bool
}

func (p * Proceso) Start() {
    go func ()  {
        for {
            p.Write = <- p.ActivoC
        }
    }()

    for p.Activo {
        if p.Write {
            fmt.Printf("id %d: %d\n", p.Id, p.I)
        }
        p.I = p.I + 1
        time.Sleep(time.Millisecond * 500)
    }
}

func (p * Proceso) Stop() {
    p.Activo = false
    close(p.ActivoC)
}
