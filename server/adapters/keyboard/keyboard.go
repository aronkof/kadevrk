package keyboard

import "github.com/bendahl/uinput"

func CreateKbd(kbdName string) (uinput.Keyboard, error) {
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte(kbdName))
	if err != nil {
		return nil, err
	}

	return keyboard, nil
}
