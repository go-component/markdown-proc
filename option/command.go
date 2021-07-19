package option

import (
	"errors"
	"path/filepath"
)

type CommandOption func(*Command) error

type Command struct {
	Output       string
	ImageDirname string
	Filename     string
}

var FilenameNotAllowEmpty = errors.New("filename not allow empty")
var OutputNotAllowEmpty = errors.New("output now allow empty")

func WithImageModeOption(imageDirname string) CommandOption {

	return func(command *Command) error {
		command.ImageDirname = imageDirname
		if imageDirname == "" {
			command.ImageDirname = filepath.Base(command.Filename)
		}
		return nil
	}
}

func WithWordModeOption() CommandOption {
	return func(command *Command) error {
		return nil
	}
}

func checkBaseOption(command *Command) error {
	if command.Filename == "" {
		return FilenameNotAllowEmpty
	}

	if command.Output == "" {
		return OutputNotAllowEmpty
	}
	return nil
}

func NewCommandOption(filename, output string, opts ...CommandOption) (*Command, error) {

	command := &Command{
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

	if err := checkBaseOption(command); err != nil{
		return nil, err
	}

	return command, nil
}
