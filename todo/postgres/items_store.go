package postgres

import (
	"database/sql"

	"github.com/edgarmijero/todo-class/todo"
	"github.com/pborman/uuid"
)

const (
	insertEventStatement   = "INSERT INTO items (id, task, completed) VALUES ($1, $2, $3)"
	findEventByIDStatement = "SELECT id, task, completed FROM items WHERE id = $1"
)

type ItemsStore struct {
	SQL *sql.DB
}

func (is ItemsStore) Insert(item *todo.Item) error {
	if item.ID == "" {
		item.ID = uuid.New()
	}
	_, err := is.SQL.Exec(insertEventStatement, item.ID, item.Task, item.Completed)
	if err != nil {
		return err
	}

	return nil
}

func (is ItemsStore) FindByID(id string) (*todo.Item, error) {
	item := &todo.Item{}

	if err := is.SQL.QueryRow(findEventByIDStatement, id).Scan(&item.ID, &item.Task, &item.Completed); err != nil {
		if err == sql.ErrNoRows {
			return nil, todo.ErrItemNotFound
		}

		return nil, err
	}

	return item, nil
}
