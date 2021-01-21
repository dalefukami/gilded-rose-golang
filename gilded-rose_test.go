package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func listWithItem(name string, sellIn int, quality int) []*Item {
	return []*Item{
		{name: name, sellIn: sellIn, quality: quality},
	}
}

func TestUpdateQuality(t *testing.T) {
	t.Run("does not change the name", func(t *testing.T) {
		items := listWithItem("foo", 0, 0)
		UpdateQuality(items)
		assert.Equal(t, "foo", items[0].name)
	})

	t.Run("quality of a normal item goes down by 1", func(t *testing.T) {
		items := listWithItem("normal", 15, 11)
		UpdateQuality(items)
		assert.Equal(t, 10, items[0].quality)
	})

	t.Run("quality of a normal item goes down by 2 when it is past it's sell by date", func(t *testing.T) {
		items := listWithItem("normal", 0, 11)
		UpdateQuality(items)
		assert.Equal(t, 9, items[0].quality)
	})

	t.Run("quality is zero don't do anything", func(t *testing.T) {
		items := listWithItem("normal", 5, 0)
		UpdateQuality(items)
		assert.Equal(t, 0, items[0].quality)
	})

	t.Run("Quality of an item is never negative", func(t *testing.T) {
		items := listWithItem("foo", 0, 0)
		UpdateQuality(items)
		assert.Equal(t, 0, items[0].quality)
	})

	t.Run("decreases the quality value by 1", func(t *testing.T) {
		items := listWithItem("foo", 5, 10)
		UpdateQuality(items)
		assert.Equal(t, 9, items[0].quality)
	})

	t.Run("decreases the quality by 2 after the sell_in has reached 0", func(t *testing.T) {
		items := listWithItem("foo", 0, 10)
		UpdateQuality(items)
		assert.Equal(t, 8, items[0].quality)
	})

	t.Run("Quality of 'Aged Brie' increases by 1 before sell_in reaches 0", func(t *testing.T) {
		items := listWithItem("Aged Brie", 5, 5)
		UpdateQuality(items)
		assert.Equal(t, 6, items[0].quality)
	})

	t.Run("Quality of 'Aged Brie' increases by 2 after sell_in reaches 0", func(t *testing.T) {
		items := listWithItem("Aged Brie", 0, 5)
		UpdateQuality(items)
		assert.Equal(t, 7, items[0].quality)
	})

	t.Run("Quality never goes above 50", func(t *testing.T) {
		items := listWithItem("Aged Brie", 0, 50)
		UpdateQuality(items)
		assert.Equal(t, 50, items[0].quality)
	})

	t.Run("'Sulfuras, Hand of Ragnaros', sell_in does not change", func(t *testing.T) {
		items := listWithItem("Sulfuras, Hand of Ragnaros", 10, 50)
		UpdateQuality(items)
		assert.Equal(t, 10, items[0].sellIn)
	})

	t.Run("'Sulfuras, Hand of Ragnaros', quality does not change", func(t *testing.T) {
		items := listWithItem("Sulfuras, Hand of Ragnaros", 10, 50)
		UpdateQuality(items)
		assert.Equal(t, 50, items[0].quality)
	})

	t.Run("Backstage passes increases in quality by 1 when sell_in is greater than 10", func(t *testing.T) {
		items := listWithItem("Backstage passes to a TAFKAL80ETC concert", 11, 10)
		UpdateQuality(items)
		assert.Equal(t, 11, items[0].quality)
	})

	t.Run("Backstage passes increases in quality by 2 when sell_in is 10 or less", func(t *testing.T) {
		items := listWithItem("Backstage passes to a TAFKAL80ETC concert", 10, 10)
		UpdateQuality(items)
		assert.Equal(t, 12, items[0].quality)
	})

	t.Run("Backstage passes increases in quality by 3 when sell_in is 5 or less", func(t *testing.T) {
		items := listWithItem("Backstage passes to a TAFKAL80ETC concert", 5, 10)
		UpdateQuality(items)
		assert.Equal(t, 13, items[0].quality)
	})

	t.Run("Backstage passes quality drops to 0 when sell_in is 0", func(t *testing.T) {
		items := listWithItem("Backstage passes to a TAFKAL80ETC concert", 0, 10)
		UpdateQuality(items)
		assert.Equal(t, 0, items[0].quality)
	})
}
