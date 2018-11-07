package api

import (
	"net/http"

	"github.com/edgarmijero/todo-class/todo"
	"github.com/labstack/echo"
)

func CreateItemsHandler(ism todo.ItemsStoreManager) echo.HandlerFunc {
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

func EditItemHandler(ism todo.ItemsStoreManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		updated_item := &todo.Item{}

		if err := c.Bind(updated_item); err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}

		// Update the item on the DB
		updated_item, err := ism.EditByID(id, updated_item)
		if err != nil {
			if err == todo.ErrItemNotFound {
				return c.String(http.StatusNotFound, "Item ("+id+") not found")
			}

			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, updated_item)
	}
}

func ShowItemHandler(ism todo.ItemsStoreManager) echo.HandlerFunc {
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

func IndexItemsHandler(ism todo.ItemsStoreManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		params := c.QueryParams()

		items, err := ism.PostgresStorage.FindByIDs(params["ids"])
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, items)
	}
}
