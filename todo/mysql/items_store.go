package mysql

import "github.com/edgarmijero/todo-class/todo"
import "log"

type ItemsStore struct {
}

func (is ItemsStore) Insert(i *todo.Item) error {
	log.Println("Mysql: ", i)

	return nil
}
