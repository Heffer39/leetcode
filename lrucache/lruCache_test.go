package lrucache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)

	assert.Equal(t, -1, obj.Get(1))
	assert.Equal(t, -1, obj.Get(2))
	assert.Equal(t, 3, obj.Get(3))
	assert.Equal(t, 4, obj.Get(4))
}

func TestLRUCacheOverwriteValue(t *testing.T) {
	obj := Constructor(3)
	obj.Put(1, 1)
	obj.Put(1, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)

	assert.Equal(t, 2, obj.Get(1))
	assert.Equal(t, 3, obj.Get(3))
	assert.Equal(t, 4, obj.Get(4))
}

func TestLRUCacheCapacity1(t *testing.T) {
	obj := Constructor(1)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)

	assert.Equal(t, -1, obj.Get(1))
	assert.Equal(t, -1, obj.Get(2))
	assert.Equal(t, -1, obj.Get(3))
	assert.Equal(t, 4, obj.Get(4))
}

func TestLRUCacheCapacity2Overwrite(t *testing.T) {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)

	assert.Equal(t, -1, obj.Get(1))
	assert.Equal(t, 3, obj.Get(2))
	assert.Equal(t, 1, obj.Get(4))
}

func TestLRUCacheCapacity2(t *testing.T) {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(3, 2)
	obj.Put(4, 3)

	assert.Equal(t, -1, obj.Get(2))
	assert.Equal(t, 2, obj.Get(3))
	assert.Equal(t, 4, obj.Get(4))
}
