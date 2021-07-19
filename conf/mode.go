package conf

import "errors"

const (
	Image = iota
	Word
)

var supportMode map[int]struct{}

func init() {
	supportMode = map[int]struct{}{
		Image: {},
		Word:  {},
	}
}

var NotMatchMode = errors.New("not match mode")

func CheckMode(mode int) error {
	if _, ok := supportMode[mode]; !ok {
		return NotMatchMode
	}

	return nil
}
