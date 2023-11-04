package clientcache

import (
	"errors"

	"github.com/arseniy96/GophKeeper/internal/client/models"
)

// Cache – объект кеша.
type Cache struct {
	mem map[int64]cacheData
}

type cacheData struct {
	data   *models.UserData
	dataID int64
}

var (
	ErrNotFound = errors.New(`not found in cache`)
)

// NewCache – функция инициализации кеша.
func NewCache() *Cache {
	return &Cache{
		mem: make(map[int64]cacheData),
	}
}

// Append – метод добавления данных в кеш.
func (c *Cache) Append(model *models.UserData) {
	c.mem[model.ID] = cacheData{
		dataID: model.ID,
		data:   model,
	}
}

// GetUserData – метод получения данных из кеша.
func (c *Cache) GetUserData(model models.UserDataModel) (*models.UserData, error) {
	d, ok := c.mem[model.ID]
	if !ok {
		return nil, ErrNotFound
	}

	return d.data, nil
}

// GetUserDataList – метод получения всех сохранённых данных (мета-данных) пользователя из кеша.
func (c *Cache) GetUserDataList() []models.UserDataList {
	records := make([]models.UserDataList, 0, len(c.mem))
	for _, d := range c.mem {
		rec := models.UserDataList{
			Name:     d.data.Name,
			DataType: d.data.DataType,
			ID:       d.data.ID,
			Version:  d.data.Version,
		}
		records = append(records, rec)
	}

	return records
}
