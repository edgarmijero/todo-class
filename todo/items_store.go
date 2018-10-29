package todo

import "errors"

var ErrItemNotFound = errors.New("Item not found")

type ItemsStore interface {
	Insert(*Item) error
	FindByID(string) (*Item, error)
}

type ItemsStoreManager struct {
	PostgresStorage ItemsStore
	MysqlStorage    ItemsStore
}

func (ism ItemsStoreManager) Insert(i *Item) error {
	if err := ism.PostgresStorage.Insert(i); err != nil {
		return err
	}

	return nil
}

func (ism ItemsStoreManager) FindByID(id string) (*Item, error) {
	return ism.PostgresStorage.FindByID(id)
}
