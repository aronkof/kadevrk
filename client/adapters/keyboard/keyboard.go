package keyboard

import (
	"sync"
	"unsafe"
)

const HC_ACTION = 0

type KeyStroke struct {
	Code  int16
	Event int16
}

func llkpFn(kb *kbListener) HOOKPROC {
	return func(nCode int, wparam WPARAM, lparam LPARAM) LRESULT {
		if nCode == HC_ACTION {
			kbdstruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))

			vkCode := int16(kbdstruct.VkCode)

			_, shouldByPass := kb.ByPassKeys[vkCode]

			kb.keyStrokes <- KeyStroke{Code: vkCode, Event: int16(wparam)}

			if shouldByPass {
				return callNextHookEx(keyboardHook, nCode, wparam, lparam)
			}
		}

		return LRESULT(uintptr(1))
	}
}

type kbListener struct {
	toggleHk   *Hotkey
	exitHk     *Hotkey
	keyStrokes chan KeyStroke
	active     bool
	hHook      HHOOK
	wg         *sync.WaitGroup
	ByPassKeys map[int16]bool
}

func NewKBListener() *kbListener {
	keyStrokes := make(chan KeyStroke)
	wg := sync.WaitGroup{}

	kb := kbListener{
		toggleHk:   &Hotkey{0, 0x76},
		exitHk:     &Hotkey{0, 0x77},
		active:     false,
		keyStrokes: keyStrokes,
		wg:         &wg,
	}

	return &kb
}

func (kb *kbListener) KeyStrokes() chan KeyStroke {
	return kb.keyStrokes
}

func (kb *kbListener) StartListener() error {
	err := registerMainHotkeys(kb.toggleHk, kb.exitHk)
	if err != nil {
		return err
	}

	kb.wg.Add(1)

	go func() {
		var msg MSG

		for getMessage(&msg, 0, 0, 0) != 0 {
			if msg.Message == WM_HOTKEY {
				if msg.WParam == exit_hk_id {
					kb.wg.Done()
					break
				}

				if msg.WParam == toggle_hk_id {
					kb.active = !kb.active

					if kb.active {
						kb.hHook = setWindowsHookEx(WH_KEYBOARD_LL, llkpFn(kb), 0, 0)
						continue
					}

					unhookWindowsHookEx(kb.hHook)
				}
			}
		}

		kb.wg.Wait()

		unhookWindowsHookEx(kb.hHook)
		close(kb.keyStrokes)
	}()

	return nil
}
