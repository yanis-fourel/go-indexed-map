package idxmap

type item[K, V any] struct {
	key K
	val V
}

type IdxMap[K comparable, V any] struct {
	s []item[K, V]
	m map[K]int
}

func New[K comparable, V any]() *IdxMap[K, V] {
	return &IdxMap[K, V]{
		s: make([]item[K, V], 0),
		m: make(map[K]int),
	}
}

func (im *IdxMap[K, V]) Append(val V) {
	var k K
	im.s = append(im.s, item[K, V]{
		key: k,
		val: val,
	})
}

func (im *IdxMap[K, V]) InsertKey(key K, val V) {
	im.s = append(im.s, item[K, V]{key, val})
	im.m[key] = len(im.s) - 1
}

func (im *IdxMap[K, V]) GetKey(key K) (res V) {
	idx, ok := im.m[key]
	if !ok {
		return
	}
	return im.s[idx].val
}

func (im *IdxMap[K, V]) GetIdx(idx int) V {
	return im.s[idx].val
}

func (im *IdxMap[K, V]) GetIdxKey(idx int) K {
	return im.s[idx].key
}

func (im *IdxMap[K, V]) RemoveIdx(idx int) {
	delete(im.m, im.s[idx].key)
	im.s = append(im.s[:idx], im.s[idx+1:]...)
	for k, i := range im.m {
		if i > idx {
			im.m[k]--
		}
	}
}

func (im *IdxMap[K, V]) RemoveKey(key K) {
	idx, ok := im.m[key]
	if !ok {
		return
	}
	im.RemoveIdx(idx)
}

func (im *IdxMap[K, V]) Iterate() []item[K, V] {
	return im.s
}

func (im *IdxMap[K, V]) Len() int {
	return len(im.s)
}
