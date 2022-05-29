package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST(
		"/whishlist",
		newAddItemHTTPHandler(
			NewAddItemToWhishlistFunc(
				newMemoryStorage(),
				newConsoleNotifier(),
			),
		),
	)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

// json over http driving adapter
func newAddItemHTTPHandler(addItem AddItemToWhishlistFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var itemToAdd ItemToAdd

		err := c.BindJSON(&itemToAdd)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		addedItem, err := addItem(c.Request.Context(), itemToAdd)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, addedItem)
	}
}

// in memory driven storage adapter
func newMemoryStorage() SaveItemFunc {
	items := make(map[int64]Item)

	return func(ctx context.Context, i Item) error {
		items[i.ID] = i

		return nil
	}
}

// terminal console driven notification adapter
func newConsoleNotifier() NotifyAboutItemAddedFunc {
	return func(ctx context.Context, item Item) error {
		fmt.Printf("> New item has been added to whishlist: [%d] %s <\n", item.ID, item.Name)

		return nil
	}
}
