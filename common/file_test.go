package common

import (
	"io"
	"testing"
)

func TestFileMd5(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileMd5(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileMd5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileMd5() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveFile(t *testing.T) {
	type args struct {
		reader  io.Reader
		dstPath string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "大文件",
			args: args{
				reader:  nil,
				dstPath: "",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SaveFile(tt.args.reader, tt.args.dstPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SaveFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
