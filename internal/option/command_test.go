package option

import (
	"github.com/go-component/markdown-proc/internal/types"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewCommandOption(t *testing.T) {
	type args struct {
		filename string
		output   string
		opts     []CommandOption
	}

	filename := "../../resource/markdown/1.md"
	output := "../../resource/markdown/output"

	fullOutput, err := filepath.Abs(output)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name    string
		args    args
		want    *types.Command
		wantErr bool
	}{
		 {
		 	name: "test-image",
		 	args: args{
				filename: filename,
				output:   output,
				opts:     []CommandOption{WithImageModeOption()},
			},
			want: &types.Command{
				Output:       fullOutput,
				ImageDirname: "1",
				Filename:     filename,
				Mode:         0,
			},
			wantErr: false,
		 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCommandOption(tt.args.filename, tt.args.output, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCommandOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommandOption() got = %v, want %v", got, tt.want)
			}
		})
	}
}
