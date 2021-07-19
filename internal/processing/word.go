package processing

import (
	"github.com/go-component/markdown-proc/types"
)

type Word struct {
	Command *types.Command
}

func (w *Word) Process() error {
	return nil
}
