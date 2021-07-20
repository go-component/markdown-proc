package internal

import (
	"github.com/go-component/markdown-proc/internal/conf"
	"github.com/go-component/markdown-proc/internal/processing"
	"github.com/go-component/markdown-proc/internal/types"
)

func Run(command *types.Command) error {

	var proc types.Processing

	switch command.Mode {
	case conf.Image:
		proc = &processing.Image{Command: command}
	case conf.Word:
		proc = &processing.Word{Command: command}
	default:
		proc = &processing.Image{Command: command}
	}

	return proc.Process()
}
