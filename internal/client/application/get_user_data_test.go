package application

import (
	"testing"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

func TestClient_GetUserData(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testClient.GetUserData(); (err != nil) != tt.wantErr {
				t.Errorf("GetUserData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_printData(t *testing.T) {
	type args struct {
		data *models.UserData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{data: testData},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := printData(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("printData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
