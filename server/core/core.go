package core

import (
	"io"
	"sync/atomic"
)

type VirtualKbd interface {
	KeyPress(key int) error
	KeyDown(key int) error
	KeyUp(key int) error
	io.Closer
}

type Dependencies struct {
	VirtualKbd VirtualKbd
}

type Rk struct {
	maxClients  int32
	currClients atomic.Int32
	kbd         VirtualKbd
}

func NewRks(d *Dependencies, maxClients int32) *Rk {
	return &Rk{kbd: d.VirtualKbd, maxClients: maxClients}
}
