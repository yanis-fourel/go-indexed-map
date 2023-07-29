package idxmap

type Item[K, V any] struct {
	Key K
	Val V
}

type IdxMap[K comparable, V any] struct {
	s []Item[K, V]
	m map[K]int
}

func New[K comparable, V any]() *IdxMap[K, V] {
	return &IdxMap[K, V]{
		s: make([]Item[K, V], 0),
		m: make(map[K]int),
	}
}

// Append new item at the end with empty key.
// Items added this way can be accessed by index only.
func (im *IdxMap[K, V]) Append(val V) {
	var k K
	im.s = append(im.s, Item[K, V]{
		Key: k,
		Val: val,
	})
}

// Overwrites existing key or append new item at the end.
// If the key is empty, this operation is equivalent to Append.
func (im *IdxMap[K, V]) InsertKey(key K, val V) {
	if _, ok := im.m[key]; ok {
		im.s[im.m[key]].Val = val
		return
	}
	var empty K
	if key == empty {
		im.Append(val)
		return
	}
	im.s = append(im.s, Item[K, V]{key, val})
	im.m[key] = len(im.s) - 1
}

// Get item value by key
// if key not found, return empty
func (im *IdxMap[K, V]) GetKey(key K) (val V) {
	idx, ok := im.m[key]
	if !ok {
		return
	}
	return im.s[idx].Val
}

// Get item index by key
// if key not found, return empty
func (im *IdxMap[K, V]) GetKeyIdx(key K) int {
	return im.m[key]
}

func (im *IdxMap[K, V]) HasKey(key K) bool {
	_, ok := im.m[key]
	return ok
}

// Get item value by index
func (im *IdxMap[K, V]) GetIdx(idx int) V {
	return im.s[idx].Val
}

// Get item key by index
func (im *IdxMap[K, V]) GetIdxKey(idx int) K {
	return im.s[idx].Key
}

// Remove item by index.
// This operation is O(n)
func (im *IdxMap[K, V]) RemoveIdx(idx int) {
	delete(im.m, im.s[idx].Key)
	im.s = append(im.s[:idx], im.s[idx+1:]...)
	for k, i := range im.m {
		if i > idx {
			im.m[k]--
		}
	}
}

// Remove item by key.
// This operation is O(n)
func (im *IdxMap[K, V]) RemoveKey(key K) {
	idx, ok := im.m[key]
	if !ok {
		return
	}
	im.RemoveIdx(idx)
}

// Will be invalidated by any modification to the IdxMap
func (im *IdxMap[K, V]) Iter() []Item[K, V] {
	return im.s
}

// The number of items in the IdxMap, including empty keys.
// This operation is O(1)
func (im *IdxMap[K, V]) Len() int {
	return len(im.s)
}

// The number of items in the IdxMap with non-empty keys.
// This operation is O(1)
func (im *IdxMap[K, V]) LenKeyed() int {
	return len(im.m)
}
