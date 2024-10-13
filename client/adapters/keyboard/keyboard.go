package keyboard

import (
	"sync"
	"unsafe"
)

const HC_ACTION = 0

type option func(*kbListener)

func WithKeysBypassList(keys []int16) option {
	return func(rkc *kbListener) {
		for _, kc := range keys {
			rkc.byPassKeys[kc] = true
		}
	}
}

type KeyStroke struct {
	Code  int16
	Event int16
}

type kbListener struct {
	toggleHk   *Hotkey
	exitHk     *Hotkey
	keyStrokes chan KeyStroke
	active     bool
	hHook      HHOOK
	wg         *sync.WaitGroup
	byPassKeys map[int16]bool
}

func llkpFn(kb *kbListener) HOOKPROC {
	return func(nCode int, wparam WPARAM, lparam LPARAM) LRESULT {
		if nCode == HC_ACTION {
			kbdstruct := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))

			vkCode := int16(kbdstruct.VkCode)

			_, shouldByPass := kb.byPassKeys[vkCode]

			kb.keyStrokes <- KeyStroke{Code: vkCode, Event: int16(wparam)}

			if shouldByPass {
				return callNextHookEx(keyboardHook, nCode, wparam, lparam)
			}
		}

		return LRESULT(uintptr(1))
	}
}

func NewKBListener(opts ...option) *kbListener {
	keyStrokes := make(chan KeyStroke)
	wg := sync.WaitGroup{}

	byPassKeys := make(map[int16]bool, 50)
	byPassKeys[toggle_hk_kc] = true
	byPassKeys[exit_hk_kc] = true

	kb := kbListener{
		toggleHk:   &Hotkey{0, toggle_hk_kc},
		exitHk:     &Hotkey{0, exit_hk_kc},
		active:     false,
		keyStrokes: keyStrokes,
		wg:         &wg,
		byPassKeys: byPassKeys,
	}

	for _, opt := range opts {
		opt(&kb)
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
