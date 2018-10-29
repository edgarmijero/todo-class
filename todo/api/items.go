package api

import (
	"net/http"

	"github.com/edgarmijero/todo-class/todo"
	"github.com/labstack/echo"
)

func CreateItems(ism todo.ItemsStoreManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		item := new(todo.Item) // &todo.Item{}

		if err := c.Bind(item); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}

		if err := ism.Insert(item); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, item)
	}
}

func ShowItems(ism todo.ItemsStoreManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Take the ID from the URL
		id := c.Param("id")
		// Fetch the item from the DB
		item, err := ism.FindByID(id)
		if err != nil {
			if err == todo.ErrItemNotFound {
				return c.String(http.StatusNotFound, "Item ("+id+") not found")
			}

			return c.String(http.StatusInternalServerError, err.Error())
		}
		// Return JSON item
		return c.JSON(http.StatusOK, item)

	}
}
