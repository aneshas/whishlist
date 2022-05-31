// A suite of ACCEPTANCE LEVEL tests against your core (hexagon)
package main

import (
	"context"
	"testing"
)

// This test is the first user of our system and also
// the first driving adapter.
// This is obviously an incomplete test but it does demark our hexagon boundary
func TestShould_Add_An_Item_To_Whishlist(t *testing.T) {
	fakeSave := func(ctx context.Context, i Item) error {
		return nil
	}

	fakeNotifiy := func(ctx context.Context, i Item) error {
		return nil
	}

	addItem := NewAddItemToWhishlistFunc(fakeSave, fakeNotifiy)

	addedItem, err := addItem(context.TODO(), ItemToAdd{
		Name: "A fancy AMZN item",
		Note: "We wants it",
	})

	if err != nil {
		t.Fatalf("no error expected, got: %v", err)
	}

	if addedItem.ItemID == 0 {
		t.Fatal("no id assigned")
	}
}

// Other acceptance level tests
