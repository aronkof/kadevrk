package keyboard

import (
	"fmt"

	"github.com/aronkof/kadev-rk/core"
	"github.com/aronkof/kadev-rk/foundation"
	"github.com/bendahl/uinput"
)

func win10Translate(code int) (core.KeyStroke, error) {
	switch code {
	// ---- modifiers ---- //
	case 91:
		return core.KeyStroke{Code: uinput.KeyLeftmeta, IsModifier: true}, nil
	case 160:
		return core.KeyStroke{Code: uinput.KeyLeftshift, IsModifier: true}, nil
	case 162:
		return core.KeyStroke{Code: uinput.KeyLeftctrl, IsModifier: true}, nil
	case 164:
		return core.KeyStroke{Code: uinput.KeyLeftalt, IsModifier: true}, nil
	case 260:
		return core.KeyStroke{Code: uinput.KeyLeftalt, IsModifier: true}, nil
	// ---- modifiers ---- //

	// ---- non alphanumeric ---- //
	case 8:
		return core.KeyStroke{Code: uinput.KeyBackspace, IsModifier: false}, nil
	case 9:
		return core.KeyStroke{Code: uinput.KeyTab, IsModifier: false}, nil
	case 13:
		return core.KeyStroke{Code: uinput.KeyEnter, IsModifier: false}, nil
	case 20:
		return core.KeyStroke{Code: uinput.KeyCapslock, IsModifier: false}, nil
	case 27:
		return core.KeyStroke{Code: uinput.KeyEsc, IsModifier: false}, nil
	case 32:
		return core.KeyStroke{Code: uinput.KeySpace, IsModifier: false}, nil
	case 33:
		return core.KeyStroke{Code: uinput.KeyPageup, IsModifier: false}, nil
	case 34:
		return core.KeyStroke{Code: uinput.KeyPagedown, IsModifier: false}, nil
	case 35:
		return core.KeyStroke{Code: uinput.KeyEnd, IsModifier: false}, nil
	case 36:
		return core.KeyStroke{Code: uinput.KeyHome, IsModifier: false}, nil
	case 45:
		return core.KeyStroke{Code: uinput.KeyInsert, IsModifier: false}, nil
	case 46:
		return core.KeyStroke{Code: uinput.KeyDelete, IsModifier: false}, nil

	case 186:
		return core.KeyStroke{Code: uinput.KeySemicolon, IsModifier: false}, nil
	case 187:
		return core.KeyStroke{Code: uinput.KeyEqual, IsModifier: false}, nil
	case 189:
		return core.KeyStroke{Code: uinput.KeyMinus, IsModifier: false}, nil
	case 220:
		return core.KeyStroke{Code: uinput.KeyBackslash, IsModifier: false}, nil
	case 191:
		return core.KeyStroke{Code: uinput.KeySlash, IsModifier: false}, nil
	case 188:
		return core.KeyStroke{Code: uinput.KeyComma, IsModifier: false}, nil
	case 190:
		return core.KeyStroke{Code: uinput.KeyDot, IsModifier: false}, nil
	case 219:
		return core.KeyStroke{Code: uinput.KeyLeftbrace, IsModifier: false}, nil
	case 221:
		return core.KeyStroke{Code: uinput.KeyRightbrace, IsModifier: false}, nil
	case 192:
		return core.KeyStroke{Code: uinput.KeyGrave, IsModifier: false}, nil
	case 222:
		return core.KeyStroke{Code: uinput.KeyApostrophe, IsModifier: false}, nil
	// ---- non alphanumeric ---- //

	// ---- arrows ---- //
	case 37:
		return core.KeyStroke{Code: uinput.KeyLeft, IsModifier: false}, nil
	case 38:
		return core.KeyStroke{Code: uinput.KeyUp, IsModifier: false}, nil
	case 39:
		return core.KeyStroke{Code: uinput.KeyRight, IsModifier: false}, nil
	case 40:
		return core.KeyStroke{Code: uinput.KeyDown, IsModifier: false}, nil
	// ---- arrows ---- //

	// ---- alphanumeric ---- //
	case 48:
		return core.KeyStroke{Code: uinput.Key0, IsModifier: false}, nil
	case 49:
		return core.KeyStroke{Code: uinput.Key1, IsModifier: false}, nil
	case 50:
		return core.KeyStroke{Code: uinput.Key2, IsModifier: false}, nil
	case 51:
		return core.KeyStroke{Code: uinput.Key3, IsModifier: false}, nil
	case 52:
		return core.KeyStroke{Code: uinput.Key4, IsModifier: false}, nil
	case 53:
		return core.KeyStroke{Code: uinput.Key5, IsModifier: false}, nil
	case 54:
		return core.KeyStroke{Code: uinput.Key6, IsModifier: false}, nil
	case 55:
		return core.KeyStroke{Code: uinput.Key7, IsModifier: false}, nil
	case 56:
		return core.KeyStroke{Code: uinput.Key8, IsModifier: false}, nil
	case 57:
		return core.KeyStroke{Code: uinput.Key9, IsModifier: false}, nil
	case 65:
		return core.KeyStroke{Code: uinput.KeyA, IsModifier: false}, nil
	case 66:
		return core.KeyStroke{Code: uinput.KeyB, IsModifier: false}, nil
	case 67:
		return core.KeyStroke{Code: uinput.KeyC, IsModifier: false}, nil
	case 68:
		return core.KeyStroke{Code: uinput.KeyD, IsModifier: false}, nil
	case 69:
		return core.KeyStroke{Code: uinput.KeyE, IsModifier: false}, nil
	case 70:
		return core.KeyStroke{Code: uinput.KeyF, IsModifier: false}, nil
	case 71:
		return core.KeyStroke{Code: uinput.KeyG, IsModifier: false}, nil
	case 72:
		return core.KeyStroke{Code: uinput.KeyH, IsModifier: false}, nil
	case 73:
		return core.KeyStroke{Code: uinput.KeyI, IsModifier: false}, nil
	case 74:
		return core.KeyStroke{Code: uinput.KeyJ, IsModifier: false}, nil
	case 75:
		return core.KeyStroke{Code: uinput.KeyK, IsModifier: false}, nil
	case 76:
		return core.KeyStroke{Code: uinput.KeyL, IsModifier: false}, nil
	case 77:
		return core.KeyStroke{Code: uinput.KeyM, IsModifier: false}, nil
	case 78:
		return core.KeyStroke{Code: uinput.KeyN, IsModifier: false}, nil
	case 79:
		return core.KeyStroke{Code: uinput.KeyO, IsModifier: false}, nil
	case 80:
		return core.KeyStroke{Code: uinput.KeyP, IsModifier: false}, nil
	case 81:
		return core.KeyStroke{Code: uinput.KeyQ, IsModifier: false}, nil
	case 82:
		return core.KeyStroke{Code: uinput.KeyR, IsModifier: false}, nil
	case 83:
		return core.KeyStroke{Code: uinput.KeyS, IsModifier: false}, nil
	case 84:
		return core.KeyStroke{Code: uinput.KeyT, IsModifier: false}, nil
	case 85:
		return core.KeyStroke{Code: uinput.KeyU, IsModifier: false}, nil
	case 86:
		return core.KeyStroke{Code: uinput.KeyV, IsModifier: false}, nil
	case 87:
		return core.KeyStroke{Code: uinput.KeyW, IsModifier: false}, nil
	case 88:
		return core.KeyStroke{Code: uinput.KeyX, IsModifier: false}, nil
	case 89:
		return core.KeyStroke{Code: uinput.KeyY, IsModifier: false}, nil
	case 90:
		return core.KeyStroke{Code: uinput.KeyZ, IsModifier: false}, nil
	// ---- alphanumeric ---- //

	// ---- FN keys ---- //
	case 112:
		return core.KeyStroke{Code: uinput.KeyF1, IsModifier: false}, nil
	case 113:
		return core.KeyStroke{Code: uinput.KeyF2, IsModifier: false}, nil
	case 114:
		return core.KeyStroke{Code: uinput.KeyF3, IsModifier: false}, nil
	case 115:
		return core.KeyStroke{Code: uinput.KeyF4, IsModifier: false}, nil
	case 116:
		return core.KeyStroke{Code: uinput.KeyF5, IsModifier: false}, nil
	case 117:
		return core.KeyStroke{Code: uinput.KeyF6, IsModifier: false}, nil
	case 118:
		return core.KeyStroke{Code: uinput.KeyF7, IsModifier: false}, nil
	case 119:
		return core.KeyStroke{Code: uinput.KeyF8, IsModifier: false}, nil
	case 120:
		return core.KeyStroke{Code: uinput.KeyF9, IsModifier: false}, nil
	case 121:
		return core.KeyStroke{Code: uinput.KeyF10, IsModifier: false}, nil
	case 122:
		return core.KeyStroke{Code: uinput.KeyF11, IsModifier: false}, nil
	case 123:
		return core.KeyStroke{Code: uinput.KeyF12, IsModifier: false}, nil
	default:
		return core.KeyStroke{}, fmt.Errorf("%d %w", code, foundation.KeyCodeNotFoundErr)
	}
	// ---- FN keys ---- //
}
