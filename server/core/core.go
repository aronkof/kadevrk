package core

import (
	"io"
)

type VirtualKbd interface {
	KeyPress(key int) error
	KeyDown(key int) error
	KeyUp(key int) error
	io.Closer
}

type Translator interface {
	Translate(os string, code int) (KeyStroke, error)
}

type KeyStroke struct {
	Code       int
	IsModifier bool
}

type Dependencies struct {
	VirtualKbd VirtualKbd
	Translator Translator
}

type Rk struct {
	kbd        VirtualKbd
	translator Translator
}

func NewRks(d *Dependencies) *Rk {
	return &Rk{kbd: d.VirtualKbd, translator: d.Translator}
}
