package config

import "fmt"

var ErrInvalidParam = func(param string, msg string) error {
	return fmt.Errorf("[Graph] invalid param %s, err: %s", param, msg)
}
