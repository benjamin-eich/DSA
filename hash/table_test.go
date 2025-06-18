package hash

import (
	"testing"
)

func TestHashTable_SetAndGet(t *testing.T) {
	ht := NewHashTable(10)
	ht.Set("foo", "bar")
	val, ok := ht.Get("foo")
	want := "bar"
	if !ok || val != want {
		t.Errorf("Get() = %s, want %v", val, want)
	}
}

func TestHashTable_OverwriteValue(t *testing.T) {
	ht := NewHashTable(10)
	ht.Set("foo", "bar")
	ht.Set("foo", "baz")
	val, ok := ht.Get("foo")
	want := "baz"
	if !ok || val != want {
		t.Errorf("Get() = %s, want %s", val, want)
	}
}

func TestHashTable_GetNonExistentKey(t *testing.T) {
	ht := NewHashTable(10)
	_, ok := ht.Get("notfound")
	want := false
	if ok != want {
		t.Errorf("Get() = %t, want %t", ok, want)
	}
}

func TestHashTable_Delete(t *testing.T) {
	ht := NewHashTable(10)
	ht.Set("foo", "bar")
	ok := ht.Delete("foo")
	want := true
	if ok != want {
		t.Errorf("Delete() = %t, want %t", ok, want)
	}
	_, found := ht.Get("foo")
	want = false
	if found != want {
		t.Errorf("Get() = _, %t, want _, %t", found, want)
	}
}

func TestHashTable_DeleteNonExistentKey(t *testing.T) {
	ht := NewHashTable(10)
	ok := ht.Delete("notfound")
	want := false
	if ok != want {
		t.Errorf("Delete() = %t, want %t", ok, want)
	}
}
