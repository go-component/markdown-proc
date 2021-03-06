package option

import (
	"errors"
	"github.com/go-component/markdown-proc/internal/conf"
	"github.com/go-component/markdown-proc/internal/types"
	"path/filepath"
	"strings"
)

type CommandOption func(command *types.Command) error

var FilenameNotAllowEmpty = errors.New("filename not allow empty")
var OutputNotAllowEmpty = errors.New("output now allow empty")

func WithImageModeOption() CommandOption {

	return func(command *types.Command) error {
		command.ImageDirname = strings.TrimSuffix(filepath.Base(command.Filename), ".md")
		command.Mode = conf.Image
		return nil
	}
}

func WithWordModeOption() CommandOption {
	return func(command *types.Command) error {
		command.Mode = conf.Word
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

	output, err := filepath.Abs(output)
	if err != nil {
		return nil, err
	}


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
