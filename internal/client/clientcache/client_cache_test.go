package clientcache

import (
	"reflect"
	"testing"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

var testData = &models.UserData{
	Name:    "testData",
	ID:      1,
	Version: 1,
}

func TestCache_Append(t *testing.T) {
	type fields struct {
		mem map[int64]cacheData
	}
	type args struct {
		model *models.UserData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "success",
			fields: fields{mem: make(map[int64]cacheData)},
			args: args{model: &models.UserData{
				Name:     "testNam",
				DataType: "password",
				Data:     []byte("tets"),
				ID:       1,
				Version:  1,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				mem: tt.fields.mem,
			}
			c.Append(tt.args.model)
			if len(c.mem) != 1 {
				t.Errorf("Cache Append error")
			}
		})
	}
}

func TestCache_GetUserData(t *testing.T) {
	type fields struct {
		mem map[int64]cacheData
	}
	type args struct {
		model models.UserDataModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.UserData
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				mem: make(map[int64]cacheData),
			},
			args: args{
				model: models.UserDataModel{ID: 1},
			},
			want: &models.UserData{
				Name:    "testData",
				ID:      1,
				Version: 1,
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				mem: make(map[int64]cacheData),
			},
			args: args{
				model: models.UserDataModel{ID: 2},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				mem: tt.fields.mem,
			}
			c.Append(testData)
			got, err := c.GetUserData(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_GetUserDataList(t *testing.T) {
	type fields struct {
		mem map[int64]cacheData
	}
	tests := []struct {
		name   string
		fields fields
		want   []models.UserDataList
	}{
		{
			name: "empty",
			fields: fields{
				mem: make(map[int64]cacheData),
			},
			want: []models.UserDataList{},
		},
		{
			name: "success",
			fields: fields{
				mem: make(map[int64]cacheData),
			},
			want: []models.UserDataList{{
				Name:    "testData",
				ID:      1,
				Version: 1,
			}},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				mem: tt.fields.mem,
			}
			if i > 0 {
				c.Append(testData)
			}

			if got := c.GetUserDataList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserDataList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCache(t *testing.T) {
	tests := []struct {
		name string
		want *Cache
	}{
		{
			name: "success",
			want: &Cache{
				mem: make(map[int64]cacheData),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
