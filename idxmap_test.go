package idxmap

import (
	"fmt"
	"sort"
	"testing"
)

func assertIs(t *testing.T, im *IdxMap[string, float64], idx int, key string, val float64) {
	_val := im.At(idx)
	_key := im.GetIdxKey(idx)

	if _val != val {
		t.Errorf("GetIdx at idx %d, expected value %f, got %f\n%v", idx, val, _val, *im)
	}
	if _key != key {
		t.Errorf("GetIdxKey at idx %d, expected key '%s', got %s\n%v", idx, key, _key, *im)
	}

	haskey := im.HasKey(key)
	if haskey != (key != "") {
		t.Errorf("HasKey at key '%s', expected %t, got %t\n%v", key, haskey, !haskey, *im)
	}

	if haskey {
		_val = im.Get(key)
		_idx := im.GetKeyIdx(key)

		if _val != val {
			t.Errorf("GetKey at key '%s', expected value %f, got %f\n%v", key, val, _val, *im)
		}
		if _idx != idx {
			t.Errorf("GetKeyIdx at key '%s', expected idx %d, got %d\n%v", key, idx, _idx, *im)
		}
	} else {
		_val = im.Get(key)
		_idx := im.GetKeyIdx(key)

		if _val != 0 {
			t.Errorf("GetKey at key '%s', expected value %f, got %f\n%v", key, val, _val, *im)
		}
		if _idx != 0 {
			t.Errorf("GetKeyIdx at key '%s', expected idx %d, got %d\n%v", key, idx, _idx, *im)
		}
	}
}

func TestEverything(t *testing.T) {
	im := New[string, float64]()

	im.Append(42)
	im.Set("pi", 3.14)
	im.Set("e", 2.71828)
	im.Set("phi", 1.61803)
	im.Append(69)
	im.Append(0)
	im.Set("sqrt(2)", 1.41421)

	im.InsertIndex(3, "tau", 6.28318)
	im.InsertIndex(7, "", 2023)
	im.InsertIndex(9, "plank", 6.62607015e-34)

	assertIs(t, im, 0, "", 42)
	assertIs(t, im, 1, "pi", 3.14)
	assertIs(t, im, 2, "e", 2.71828)
	assertIs(t, im, 3, "tau", 6.28318)
	assertIs(t, im, 4, "phi", 1.61803)
	assertIs(t, im, 5, "", 69)
	assertIs(t, im, 6, "", 0)
	assertIs(t, im, 7, "", 2023)
	assertIs(t, im, 8, "sqrt(2)", 1.41421)
	assertIs(t, im, 9, "plank", 6.62607015e-34)
	if im.Len() != 10 {
		t.Errorf("Len, expected 10, got %d\n%v", im.Len(), *im)
	}
	if im.LenKeyed() != 6 {
		t.Errorf("LenKeyed, expected 6, got %d\n%v", im.LenKeyed(), *im)
	}
	fmt.Println(im)

	im.RemoveAt(0)
	im.RemoveAt(4)
	im.RemoveAt(7)
	im.Remove("pi")
	im.Remove("tau")
	im.Remove("sqrt(2)")

	assertIs(t, im, 0, "e", 2.71828)
	assertIs(t, im, 1, "phi", 1.61803)
	assertIs(t, im, 2, "", 0)
	assertIs(t, im, 3, "", 2023)
	if im.Len() != 4 {
		t.Errorf("Len, expected 4, got %d\n%v", im.Len(), *im)
	}
	if im.LenKeyed() != 2 {
		t.Errorf("LenKeyed, expected 2, got %d\n%v", im.LenKeyed(), *im)
	}
	fmt.Println(im)

	im.Set("e", 2.71828182845904523536028747135266249775724709369995)
	im.Set("phi", 1.61803398874989484820458683436563811772030917980576)
	im.Set("plank", 6.62607015e-34)
	im.SetIdx(2, 1)
	im.SetIdx(3, 2024)

	assertIs(t, im, 0, "e", 2.71828182845904523536028747135266249775724709369995)
	assertIs(t, im, 1, "phi", 1.61803398874989484820458683436563811772030917980576)
	assertIs(t, im, 2, "", 1)
	assertIs(t, im, 3, "", 2024)
	assertIs(t, im, 4, "plank", 6.62607015e-34)
	if im.Len() != 5 {
		t.Errorf("Len, expected 5, got %d\n%v", im.Len(), *im)
	}
	if im.LenKeyed() != 3 {
		t.Errorf("LenKeyed, expected 3, got %d\n%v", im.LenKeyed(), *im)
	}
	fmt.Println(im)

	im.Set("", 999)
	im.SetIdx(0, 0)
	im.SetIdxKey(0, "zero")

	assertIs(t, im, 0, "zero", 0)
	assertIs(t, im, 1, "phi", 1.61803398874989484820458683436563811772030917980576)
	assertIs(t, im, 2, "", 1)
	assertIs(t, im, 3, "", 2024)
	assertIs(t, im, 4, "plank", 6.62607015e-34)
	assertIs(t, im, 5, "", 999)
	if im.Len() != 6 {
		t.Errorf("Len, expected 6, got %d\n%v", im.Len(), *im)
	}
	if im.LenKeyed() != 3 {
		t.Errorf("LenKeyed, expected 3, got %d\n%v", im.LenKeyed(), *im)
	}
	fmt.Println(im)

	s := im.Slice()
	sort.Slice(s, func(i, j int) bool { return s[i].Val < s[j].Val })
	im = From(s)

	assertIs(t, im, 0, "zero", 0)
	assertIs(t, im, 1, "plank", 6.62607015e-34)
	assertIs(t, im, 2, "", 1)
	assertIs(t, im, 3, "phi", 1.61803398874989484820458683436563811772030917980576)
	assertIs(t, im, 4, "", 999)
	assertIs(t, im, 5, "", 2024)

	s = im.Slice()
	sort.SliceStable(s, func(i, j int) bool { return s[i].Key < s[j].Key })
	im = From(s)

	assertIs(t, im, 0, "", 1)
	assertIs(t, im, 1, "", 999)
	assertIs(t, im, 2, "", 2024)
	assertIs(t, im, 3, "phi", 1.61803398874989484820458683436563811772030917980576)
	assertIs(t, im, 4, "plank", 6.62607015e-34)
	assertIs(t, im, 5, "zero", 0)
}
