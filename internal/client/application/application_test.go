package application

import (
	"testing"
	"time"

	"github.com/arseniy96/GophKeeper/internal/client/config"
)

func TestClient_GetAuthToken(t *testing.T) {
	type fields struct {
		AuthToken string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty token",
			fields: fields{
				AuthToken: "",
			},
			want: "",
		},
		{
			name: "token is present",
			fields: fields{
				AuthToken: "some_token",
			},
			want: "some_token",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				AuthToken: tt.fields.AuthToken,
			}
			if got := c.GetAuthToken(); got != tt.want {
				t.Errorf("GetAuthToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func _TestClient_GetTimeout(t *testing.T) {
	type fields struct {
		Config    *config.Config
		AuthToken string
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "success",
			fields: fields{
				Config: &config.Config{ConnectionTimeout: 5},
			},
			want: time.Duration(5 * time.Second),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Config:    tt.fields.Config,
				AuthToken: tt.fields.AuthToken,
			}
			if got := c.GetTimeout(); got != tt.want {
				t.Errorf("GetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateAuthToken(t *testing.T) {
	type fields struct {
		AuthToken string
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "updated",
			fields: fields{AuthToken: ""},
			args:   args{token: "some"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				AuthToken: tt.fields.AuthToken,
			}
			c.UpdateAuthToken(tt.args.token)
			if c.AuthToken != tt.args.token {
				t.Errorf("UpdateAuthToken error")
			}
		})
	}
}
