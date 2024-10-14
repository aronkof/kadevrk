package core

import (
	"errors"

	"github.com/aronkof/kadev-rk/foundation"
)

func (rk *Rk) DispatchKeyEvent(os string, code int, keyDown bool) error {
	keyStroke, err := rk.translator.Translate(os, code)
	if err != nil {
		return err
	}

	if errors.Is(err, foundation.KeyCodeNotFoundErr) {
		return nil
	}

	if keyDown {
		return rk.kbd.KeyDown(keyStroke.Code)
	}

	return rk.kbd.KeyUp(keyStroke.Code)
}
