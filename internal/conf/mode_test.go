package conf

import "testing"

func TestCheckMode(t *testing.T) {
	type args struct {
		mode int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test image",
			args: args{mode: 0},
			wantErr: false,
		},
		{
			name: "test word",
			args: args{mode: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckMode(tt.args.mode); (err != nil) != tt.wantErr {
				t.Errorf("CheckMode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
