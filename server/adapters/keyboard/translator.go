package keyboard

import (
	"fmt"

	"github.com/aronkof/kadev-rk/core"
	"github.com/aronkof/kadev-rk/foundation"
)

type Translator struct{}

func (*Translator) Translate(os string, code int) (core.KeyStroke, error) {
	switch os {
	case "windows-11", "win-11", "windows11", "windows 11":
		return core.KeyStroke{}, fmt.Errorf("%s %w", os, foundation.OsNotSupported)

	case "windows-10", "win-10", "windows10", "windows 10":
		return win10Translate(code)

	default:
		return core.KeyStroke{}, fmt.Errorf("unknown OS [%s]", os)
	}
}

func NewTranslator() *Translator {
	return &Translator{}
}
