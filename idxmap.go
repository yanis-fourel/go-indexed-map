package idxmap

type item struct {
	key int
	val int
}

type IdxMap struct {
	s []item
	m map[int]int
}

func New() *IdxMap {
	return &IdxMap{
		s: make([]item, 0),
		m: make(map[int]int),
	}
}

func (im *IdxMap) Append(val int) {
	im.s = append(im.s, item{0, val})
}

func (im *IdxMap) InsertKey(key int, val int) {
	im.s = append(im.s, item{key, val})
	im.m[key] = len(im.s) - 1
}

func (im *IdxMap) GetKey(key int) int {
	idx, ok := im.m[key]
	if !ok {
		return 0
	}
	return im.s[idx].val
}

func (im *IdxMap) GetIdx(idx int) int {
	return im.s[idx].val
}

func (im *IdxMap) RemoveIdx(idx int) {
	delete(im.m, im.s[idx].key)
	im.s = append(im.s[:idx], im.s[idx+1:]...)
	for k, i := range im.m {
		if i > idx {
			im.m[k]--
		}
	}
}

func (im *IdxMap) RemoveKey(key int) {
	idx, ok := im.m[key]
	if !ok {
		return
	}
	im.RemoveIdx(idx)
}

func (im *IdxMap) Iterate() []item {
	return im.s
}

func (im *IdxMap) Len() int {
	return len(im.s)
}
