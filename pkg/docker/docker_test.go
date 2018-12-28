package docker

import (
	"testing"
)

func TestGetExposePorts(t *testing.T) {
	type args struct {
		image string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"no expose port", args{"busybox"}, 0},
		{"expose port", args{"pythonbox"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetExposePorts(tt.args.image); len(got) != tt.want {
				t.Errorf("GetExposePorts() = %v, want %v", got, tt.want)
			}
		})
	}
}
