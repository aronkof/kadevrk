package foundation

import "errors"

var (
	KeyCodeNotFoundErr = errors.New("key code not found")
	OsNotSupported     = errors.New("OS not supported yet")
)
