package factory

import "testing"

func TestInitConfigFactory(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				path: ".",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitConfigFactory(); (err != nil) != tt.wantErr {
				t.Errorf("InitConfigFactory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
