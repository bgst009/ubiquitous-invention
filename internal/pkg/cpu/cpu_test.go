package cpu

import "testing"

func TestGetUsageByPID(t *testing.T) {
	type args struct {
		pid int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
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
			got := GetUsageByPID(tt.args.pid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsageByPID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUsageByPID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
