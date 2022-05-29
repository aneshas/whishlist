// This is my hexagon. I don't care what's outside, in this folder or any
// other for that matter. I don't care where this code is phisically located
// as long as it is decoupled from any driving or driven adapter be it mock,
// fake or any technology specific "real" implementation.
package main

import (
	"context"
	"time"
)

// Item represents whishlist item. It could be your domain model, ddd aggregate etc...
type Item struct {
	ID      int64
	Name    string
	Note    string
	AddedOn time.Time
}

// ItemToAdd represents item to be added to whishlist
type ItemToAdd struct {
	Name string
	Note string
}

// AddedItem represents an item that has been added to whishlist
type AddedItem struct {
	ItemID int64
}

// AddItemToWhishlistFunc represents add item to whishlist use case to be
// used by any driving adapter.
// NOTE:
// - I didn't even need to define it explicitly, driven adapter can have direct
//   dependency on the implementation
// - Words like "adapter" and "port" usually have no place in the hexagon
//   I prefer domain/problem related language and so should you.
type AddItemToWhishlistFunc func(context.Context, ItemToAdd) (*AddedItem, error)

// SaveItemFunc represents item saving mechanism
type SaveItemFunc func(context.Context, Item) error

// NotifyAboutItemAddedFunc represents example notification mechanism
type NotifyAboutItemAddedFunc func(context.Context, Item) error

// NewAddItemToWhishlistFunc constructs item adding use case
func NewAddItemToWhishlistFunc(save SaveItemFunc, notify NotifyAboutItemAddedFunc) AddItemToWhishlistFunc {
	return func(ctx context.Context, i ItemToAdd) (*AddedItem, error) {
		item := Item{
			ID:      time.Now().UnixMilli(),
			Name:    i.Name,
			Note:    i.Note,
			AddedOn: time.Now(),
		}

		err := save(ctx, item)
		if err != nil {
			return nil, err
		}

		err = notify(ctx, item)
		if err != nil {
			return nil, err
		}

		return &AddedItem{
			ItemID: item.ID,
		}, nil
	}
}
