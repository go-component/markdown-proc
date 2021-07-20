package fileutil

import "testing"

func TestExt(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ext-test1",
			args: args{path: "test1/test2/base.txt"},
			want: ".txt",
		},
		{
			name: "ext-test2",
			args: args{path: "test1/test2/base.txt!extra1"},
			want: ".txt",
		},
		{
			name: "ext-test3",
			args: args{path: "test1/test2/base.txt?extra2"},
			want: ".txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ext(tt.args.path); got != tt.want {
				t.Errorf("Ext() = %v, want %v", got, tt.want)
			}
		})
	}
}
