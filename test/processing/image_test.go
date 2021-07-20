package processing

import (
	"github.com/go-component/markdown-proc/internal/processing"
	"github.com/go-component/markdown-proc/internal/types"
	"github.com/go-component/markdown-proc/option"
	"testing"
)

func TestImage_Process(t *testing.T) {

	type fields struct {
		Command *types.Command
	}

	commandOption, err := option.NewCommandOption(
		"../../resource/markdown/1.md",
		"../../resource/markdown/output",
		option.WithImageModeOption(),
	)
	if err != nil{
		panic(err)
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "process-test1",
			fields: fields{Command: commandOption},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &processing.Image{
				Command: tt.fields.Command,
			}
			if err := i.Process(); (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
