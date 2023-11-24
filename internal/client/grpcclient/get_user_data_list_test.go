package grpcclient

import (
	"reflect"
	"testing"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

func TestClient_GetUserDataList(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.UserDataList
		wantErr bool
	}{
		{
			name: "success",
			want: []models.UserDataList{{
				ID:       1,
				Name:     "testData",
				DataType: "password",
				Version:  1,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testClient.GetUserDataList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserDataList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserDataList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
