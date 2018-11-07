package postgres

import (
	"database/sql"

	"github.com/edgarmijero/todo-class/todo"
	"github.com/lib/pq"
	"github.com/pborman/uuid"
)

const (
	insertStatement    = "INSERT INTO items (id, task, completed) VALUES ($1, $2, $3)"
	findByIDStatement  = "SELECT id, task, completed FROM items WHERE id = $1"
	findByIDsStatement = "SELECT id, task, completed FROM items WHERE id = ANY ($1)"
	updateStatement    = "UPDATE items SET task = $1, completed = $2 WHERE id = $3"
)

type ItemsStore struct {
	SQL *sql.DB
}

func (is ItemsStore) FindByIDs(ids []string) ([]*todo.Item, error) {
	rows, err := is.SQL.Query(findByIDsStatement, pq.Array(ids))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return is.sqlRowsToItems(rows)
}

func (is ItemsStore) sqlRowsToItems(rows *sql.Rows) ([]*todo.Item, error) {
	items := []*todo.Item{}

	for rows.Next() {
		item := new(todo.Item)

		if err := rows.Scan(&item.ID, &item.Task, &item.Completed); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (is ItemsStore) Insert(item *todo.Item) error {
	if item.ID == "" {
		item.ID = uuid.New()
	}
	_, err := is.SQL.Exec(insertStatement, item.ID, item.Task, item.Completed)
	if err != nil {
		return err
	}

	return nil
}

func (is ItemsStore) FindByID(id string) (*todo.Item, error) {
	item := &todo.Item{}

	if err := is.SQL.QueryRow(findByIDStatement, id).Scan(&item.ID, &item.Task, &item.Completed); err != nil {
		if err == sql.ErrNoRows {
			return nil, todo.ErrItemNotFound
		}

		return nil, err
	}

	return item, nil
}

func (is ItemsStore) EditByID(id string, item *todo.Item) (*todo.Item, error) {
	updated_item, err := is.FindByID(id)
	if err != nil {
		return nil, err
	}

	if item.Task != "" {
		updated_item.Task = item.Task
	}

	_, err = is.SQL.Exec(updateStatement, &updated_item.Task, &updated_item.Completed, updated_item.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, todo.ErrItemNotFound
		}

		return nil, err
	}

	return updated_item, nil
}
