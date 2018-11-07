package main

import (
	"database/sql"
	"log"

	"github.com/edgarmijero/todo-class/todo"
	"github.com/edgarmijero/todo-class/todo/api"
	"github.com/edgarmijero/todo-class/todo/postgres"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:@127.0.0.1/todo_dev?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	postItemsStore := postgres.ItemsStore{
		SQL: db,
	}

	itemsStoreManager := todo.ItemsStoreManager{
		PostgresStorage: postItemsStore,
	}

	e := echo.New()

	e.GET("/healthz", api.Healthz)
	e.POST("/items", api.CreateItemsHandler(itemsStoreManager))
	e.GET("/items/:id", api.ShowItemHandler(itemsStoreManager))
	e.PUT("/items/:id", api.EditItemHandler(itemsStoreManager))
	e.GET("/items/", api.IndexItemsHandler(itemsStoreManager))

	e.Start(":8080")
}
