package common

import (
	"testing"
)

func TestExexCmd(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test-cpu",
			args: args{
				cmd: GetCmdByPID(5555, "cpu"),
			},
		}, {
			name: "test-mem",
			args: args{
				cmd: GetCmdByPID(5555, "mem"),
			},
		}, {
			name: "test-name",
			args: args{
				cmd: GetCmdByPID(5555, "name"),
			},
		}, {
			name: "test-default",
			args: args{
				cmd: GetCmdByPID(5555, "s"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExexCmd(tt.args.cmd)
		})
	}
}

func TestGetCmdByPID(t *testing.T) {
	type args struct {
		pid int
		w   string
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
				w:   "cpu",
			},
			want: `top -bn 1 -p 5555| tail -1 | awk '{ssd=NF-3} {print $ssd }'`,
		},
		{
			name: "test-mem",
			args: args{
				pid: 5555,
				w:   "mem",
			},
			want: `top -bn 1 -p 5555| tail -1 | awk '{ssd=NF-6} {print $ssd }'`,
		},
		{
			name: "test-name",
			args: args{
				pid: 5555,
				w:   "name",
			},
			want: `top -bn 1 -p 5555| tail -1 | awk '{ssd=NF-0} {print $ssd }'`,
		},
		{
			name: "test-default",
			args: args{
				pid: 5555,
				w:   "s",
			},
			want: `top -bn 1 -p 5555`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCmdByPID(tt.args.pid, tt.args.w); got != tt.want {
				t.Errorf("GetCmdByPID() = %v, want %v", got, tt.want)
			}
		})
	}
}
