package main

import (
	"context"
	"time"
)

type Item struct {
	Name    string
	Note    string
	AddedOn time.Time
}

type ItemToAdd struct {
	Name string
	Note string
}

type AddItemToWhishlistFunc func(context.Context, ItemToAdd) error

type SaveItemFunc func(context.Context, Item) error

type NotifyAboutItemAddedFunc func(context.Context, string) error

func NewAddItemToWhishlistFunc(save SaveItemFunc, notify NotifyAboutItemAddedFunc) AddItemToWhishlistFunc {
	return func(ctx context.Context, i ItemToAdd) error {
		item := Item{
			Name:    i.Name,
			Note:    i.Note,
			AddedOn: time.Now(),
		}

		err := save(ctx, item)
		if err != nil {
			return err
		}

		return notify(ctx, item.Name)
	}
}
