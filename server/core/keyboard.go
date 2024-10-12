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

	if keyDown && keyStroke.IsModifier {
		if keyStroke.IsModifier {
			return rk.kbd.KeyDown(keyStroke.Code)
		} else {
			return rk.kbd.KeyPress(keyStroke.Code)
		}
	}

	return rk.kbd.KeyUp(keyStroke.Code)
}
