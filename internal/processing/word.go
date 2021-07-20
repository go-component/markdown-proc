package processing

import (
	"github.com/go-component/markdown-proc/internal/types"
)

type Word struct {
	Command *types.Command
}

func (w *Word) Process() error {
	return nil
}
