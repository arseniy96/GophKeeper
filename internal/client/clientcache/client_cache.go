package clientcache

import (
	"errors"
	"sync"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

// Cache – объект кеша.
type Cache struct {
	mem *sync.Map
}

var (
	ErrNotFound = errors.New(`not found in cache`)
)

// NewCache – функция инициализации кеша.
func NewCache() *Cache {
	return &Cache{
		mem: &sync.Map{},
	}
}

// Append – метод добавления данных в кеш.
func (c *Cache) Append(model *models.UserData) {
	c.mem.Store(model.ID, model)
}

// GetUserData – метод получения данных из кеша.
func (c *Cache) GetUserData(model models.UserDataModel) (*models.UserData, error) {
	d, ok := c.mem.Load(model.ID)
	if !ok {
		return nil, ErrNotFound
	}

	return d.(*models.UserData), nil
}

// GetUserDataList – метод получения всех сохранённых данных (мета-данных) пользователя из кеша.
func (c *Cache) GetUserDataList() []models.UserDataList {
	records := make([]models.UserDataList, 0)
	c.mem.Range(func(k, v interface{}) bool {
		rec := models.UserDataList{
			Name:     v.(*models.UserData).Name,
			DataType: v.(*models.UserData).DataType,
			ID:       k.(int64),
			Version:  v.(*models.UserData).Version,
		}
		records = append(records, rec)

		return true
	})

	return records
}
