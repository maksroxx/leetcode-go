package lrucache

import (
	"testing"
)

func TestLRUCache_Basic(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if val := cache.Get(1); val != 1 {
		t.Errorf("Expected Get(1) == 1, got %d", val)
	}

	cache.Put(3, 3)

	if val := cache.Get(2); val != -1 {
		t.Errorf("Expected Get(2) == -1 (evicted), got %d", val)
	}

	cache.Put(4, 4)

	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected Get(1) == -1 (evicted), got %d", val)
	}

	if val := cache.Get(3); val != 3 {
		t.Errorf("Expected Get(3) == 3, got %d", val)
	}

	if val := cache.Get(4); val != 4 {
		t.Errorf("Expected Get(4) == 4, got %d", val)
	}
}

func TestLRUCache_UpdateValue(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 10)

	if val := cache.Get(1); val != 10 {
		t.Errorf("Expected updated Get(1) == 10, got %d", val)
	}
}

func TestLRUCache_AccessMovesToFront(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)

	if val := cache.Get(2); val != -1 {
		t.Errorf("Expected Get(2) == -1 (evicted), got %d", val)
	}
}
