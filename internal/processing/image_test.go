package processing

import (
	"github.com/go-component/markdown-proc/internal/types"
	"path/filepath"
	"strings"
	"testing"
)

func TestImage_Process(t *testing.T) {

	type fields struct {
		Command *types.Command
	}

	output := "../../resource/markdown/output"

	output, err := filepath.Abs(output)
	if err != nil {
		panic(err)
	}

	filename := "../../resource/markdown/1.md"

	command := &types.Command{
		Output:       output,
		Filename:     filename,
		ImageDirname: strings.TrimSuffix(filepath.Base(filename), ".md"),
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "process-test1",
			fields:  fields{Command: command},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Image{
				Command: tt.fields.Command,
			}
			if err := i.Process(); (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
