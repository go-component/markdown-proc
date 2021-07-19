package option

import (
	"errors"
	"github.com/go-component/markdown-proc/internal/processing"
	"github.com/go-component/markdown-proc/types"
	"path/filepath"
	"strings"
)

type CommandOption func(command *types.Command) error

var FilenameNotAllowEmpty = errors.New("filename not allow empty")
var OutputNotAllowEmpty = errors.New("output now allow empty")

func WithImageModeOption() CommandOption {

	return func(command *types.Command) error {
		command.ImageDirname = strings.Trim(filepath.Base(command.Filename), ".md")
		command.Processing = &processing.Image{
			Command: command,
		}
		return nil
	}
}

func WithWordModeOption() CommandOption {
	return func(command *types.Command) error {
		command.Processing = &processing.Word{
			Command: command,
		}
		return nil
	}
}

func checkBaseOption(command *types.Command) error {
	if command.Filename == "" {
		return FilenameNotAllowEmpty
	}

	if command.Output == "" {
		return OutputNotAllowEmpty
	}
	return nil
}

func NewCommandOption(filename, output string, opts ...CommandOption) (*types.Command, error) {

	command := &types.Command{
		Output:       output,
		ImageDirname: "",
		Filename:     filename,
	}

	for _, opt := range opts {
		err := opt(command)

		if err != nil {
			return nil, err
		}
	}

	if err := checkBaseOption(command); err != nil {
		return nil, err
	}

	return command, nil
}
