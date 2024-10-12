package keyboard

import (
	"github.com/aronkof/kadev-rk/core"
	"github.com/bendahl/uinput"
)

func win10Translate(code int) (core.KeyStroke, error) {
	switch code {
	// ---- modifiers ---- //
	case 160:
		return core.KeyStroke{Code: uinput.KeyLeftshift, IsModifier: true}, nil
	case 162:
		return core.KeyStroke{Code: uinput.KeyLeftctrl, IsModifier: true}, nil
	case 164:
		return core.KeyStroke{Code: uinput.KeyLeftalt, IsModifier: true}, nil
	case 91:
		return core.KeyStroke{Code: uinput.KeyLeftmeta, IsModifier: false}, nil
	// ---- modifiers ---- //

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
	case 37:
		return core.KeyStroke{Code: uinput.KeyLeft, IsModifier: false}, nil
	case 38:
		return core.KeyStroke{Code: uinput.KeyUp, IsModifier: false}, nil
	case 39:
		return core.KeyStroke{Code: uinput.KeyRight, IsModifier: false}, nil
	case 40:
		return core.KeyStroke{Code: uinput.KeyDown, IsModifier: false}, nil
	case 45:
		return core.KeyStroke{Code: uinput.KeyInsert, IsModifier: false}, nil
	case 46:
		return core.KeyStroke{Code: uinput.KeyDelete, IsModifier: false}, nil
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
	}

	// VK_MULTIPLY	106	Multiply key
	// VK_ADD	107	Add key
	// VK_SEPARATOR	108	Separator key
	// VK_SUBTRACT	109	Subtract key
	// VK_DECIMAL	110	Decimal key
	// VK_DIVIDE	111	Divide key

	// VK_F1	112	F1 key
	// VK_F2	113	F2 key
	// VK_F3	114	F3 key
	// VK_F4	115	F4 key
	// VK_F5	116	F5 key
	// VK_F6	117	F6 key
	// VK_F7	118	F7 key
	// VK_F8	119	F8 key
	// VK_F9	120	F9 key
	// VK_F10	121	F10 key
	// VK_F11	122	F11 key
	// VK_F12	123	F12 key
	// VK_F13	124	F13 key
	// VK_F24	135	F24 key

	// VK_OEM_1	186	Used for miscellaneous characters; varies by keyboard. For the US standard keyboard, the ;: key
	// VK_OEM_PLUS	187	For any country/region, the + key
	// VK_OEM_COMMA	188	For any country/region, the , key
	// VK_OEM_MINUS	189	For any country/region, the - key
	// VK_OEM_PERIOD	190	For any country/region, the . key
	// VK_OEM_2	191	Used for miscellaneous characters; varies by keyboard. For the US standard keyboard, the /? key
	// VK_OEM_3	192	Used for miscellaneous characters; varies by keyboard. For the US standard keyboard, the `~ key
	// VK_OEM_4	219	Used for miscellaneous characters; varies by keyboard. For the US standard keyboard, the [ { key
	// VK_OEM_5	220	Used for miscellaneous characters; varies by keyboard. For the US standard keyboard, the | key
	// VK_OEM_6	221	Used for miscellaneous characters; varies by keyboard. For the US standard keyboard, the ] } key
	// VK_OEM_7	222	Used for miscellaneous characters; varies by keyboard. For the US standard keyboard, the ' " key
	// VK_OEM_8	223	Used for miscellaneous characters; varies by keyboard.

	return core.KeyStroke{}, nil
}
