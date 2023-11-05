package application

import (
	"net"
	"testing"

	"github.com/arseniy96/GophKeeper/internal/client/config"
	"github.com/arseniy96/GophKeeper/src/logger"
)

func TestNewClient(t *testing.T) {
	type args struct {
		l *logger.Logger
		c *config.Config
	}
	tests := []struct {
		name    string
		args    args
		grpcRun bool
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				l: &logger.Logger{},
				c: &config.Config{
					ChanSize:          1,
					ConnectionTimeout: 5,
					Host:              "localhost:3200",
				},
			},
			grpcRun: true,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				l: &logger.Logger{},
				c: &config.Config{},
			},
			grpcRun: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.grpcRun {
				conn, err := net.Listen("tcp", ":3200")
				if err != nil {
					t.Errorf("gRPC server start error = %v", err)
				}
				defer func() {
					_ = conn.Close()
				}()
			}

			_, err := NewClient(tt.args.l, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
