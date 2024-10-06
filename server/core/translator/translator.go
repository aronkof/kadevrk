package translator

import "fmt"

func Translate(os string, code int) (int, error) {
	switch os {
	case "windows-11", "win-11", "windows11", "windows 11":
		return win11Translate(code)

	case "windows-10", "win-10", "windows10", "windows 10":
		return win10Translate(code)

	default:
		return 0, fmt.Errorf("unknown OS [%s]", os)
	}
}
