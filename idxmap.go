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

// Create IdxMap from slice of items.
// The provided slice will not be copied but used as is.
// Any subsequent changes to the slice will invalidate the IdxMap.
func From[K comparable, V any](s []Item[K, V]) *IdxMap[K, V] {
	var emptykey K
	m := make(map[K]int)
	for idx, item := range s {
		if item.Key != emptykey {
			m[item.Key] = idx
		}
	}
	return &IdxMap[K, V]{s, m}
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

// Append new item at index.
// If key already exists, it will be overwritten instead, and `idx` will be
// ignored.
func (im *IdxMap[K, V]) InsertIndex(idx int, key K, val V) {
	if _, ok := im.m[key]; ok {
		im.s[im.m[key]].Val = val
		return
	}
	im.s = append(im.s[:idx+1], im.s[idx:]...)
	im.s[idx] = Item[K, V]{key, val}
	for k, i := range im.m {
		if i >= idx {
			im.m[k]++
		}
	}
	var empty K
	if key != empty {
		im.m[key] = idx
	}
}

// Overwrites existing key or append new item at the end.
// If the key is empty, this operation is equivalent to Append.
func (im *IdxMap[K, V]) Set(key K, val V) {
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

// Set item value by index
func (im *IdxMap[K, V]) SetIdx(idx int, val V) {
	im.s[idx].Val = val
}

// Set item key by index
func (im *IdxMap[K, V]) SetIdxKey(idx int, key K) {
	var emptykey K
	delete(im.m, im.s[idx].Key)
	im.s[idx].Key = key

	if key != emptykey {
		im.m[key] = idx
	}
}

// Get item value by key
// if key not found, return empty
func (im *IdxMap[K, V]) Get(key K) (val V) {
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
func (im *IdxMap[K, V]) At(idx int) V {
	return im.s[idx].Val
}

// Get item key by index
func (im *IdxMap[K, V]) GetIdxKey(idx int) K {
	return im.s[idx].Key
}

// Remove item by index.
// This operation is O(n)
func (im *IdxMap[K, V]) RemoveAt(idx int) {
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
func (im *IdxMap[K, V]) Remove(key K) {
	idx, ok := im.m[key]
	if !ok {
		return
	}
	im.RemoveAt(idx)
}

// Any modification of the slice will invalidate the IdxMap, and any
// modification of the IdxMap will invalidate the slice.
//
// This can be used in combination with `From` to apply any arbitrary
// modifications to the slice and then rebuild the IdxMap with zero copies
// of the slice content.
//
// eg:
// ```
//
//	func sortIdxMap(im *IdxMap[string, float64]) {
//	    s := im.Slice()
//	    sort.Slice(s, func(i, j int) bool { return s[i].Val < s[j].Val })
//	    im = From(s)
//	}
//
// ```
func (im *IdxMap[K, V]) Slice() []Item[K, V] {
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
