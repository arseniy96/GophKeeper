package application

import "testing"

func TestClient_GetUserDataList(t *testing.T) {
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
			if err := testClient.GetUserDataList(); (err != nil) != tt.wantErr {
				t.Errorf("GetUserDataList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
