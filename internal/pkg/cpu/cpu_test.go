package cpu

import "testing"

func TestGetUsageByPID(t *testing.T) {
	type args struct {
		pid int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-cpu",
			args: args{
				pid: 5555,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsageByPID(tt.args.pid); got != tt.want {
				t.Errorf("GetUsageByPID() = %v, want %v", got, tt.want)
			}
		})
	}
}
