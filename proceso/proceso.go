package proceso

import "fmt"
type Proceso struct {
    Id uint64
    Activo bool
}

func (p * Proceso) Start(c chan uint64) {
    for p.Activo {
        c <- p.Id
    }
}

func (p * Proceso) Stop() {
    p.Activo = false
}
