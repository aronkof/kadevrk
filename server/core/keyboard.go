package core

import "github.com/aronkof/kadev-rk/core/translator"

func (s *Rk) DispatchKeyEvent(os string, code int, keyDown bool) error {
	keyCode, err := translator.Translate(os, code)
	if err != nil {
		return err
	}

	if keyDown {
		return s.kbd.KeyDown(keyCode)
	}

	return s.kbd.KeyUp(keyCode)
}
