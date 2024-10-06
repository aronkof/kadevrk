package keyboard

import "fmt"

type Hotkey struct {
	Modifiers int
	KeyCode   int
}

const (
	toggle_hk_id = 101
	exit_hk_id   = 102
)

const (
	ModAlt = 1 << iota
	ModCtrl
	ModShift
	ModWin
)

func hkName(id int16) string {
	if id == toggle_hk_id {
		return "TOGGLE"
	}

	if id == exit_hk_id {
		return "EXIT"
	}

	return "UNKNOWN"
}

func registerMainHotkeys(toggle, exit *Hotkey) error {
	err := registerHotKey(toggle, toggle_hk_id)
	if err != nil {
		return err
	}

	err = registerHotKey(exit, exit_hk_id)
	if err != nil {
		return err
	}

	return nil
}

func registerHotKey(hk *Hotkey, id int16) error {
	r1, _, _ := reghotkey.Call(0, uintptr(id), uintptr(hk.Modifiers), uintptr(hk.KeyCode))

	if r1 != 1 {
		return fmt.Errorf("could not register %s hotkey for keycode %d", hkName(id), hk.KeyCode)
	}

	return nil
}
