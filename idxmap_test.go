package idxmap

import (
	"fmt"
	"testing"
)

func TestEverything(t *testing.T) {
	im := New[int, int]()

	im.Append(1)
	im.Append(2)
	im.Append(3)
	im.InsertKey(42, 4)
	im.Append(5)
	fmt.Println(im)

	if im.Len() != 5 {
		t.Errorf("Expected len 5, got %d", im.Len())
	}

	if im.GetIdx(0) != 1 {
		t.Errorf("Index 0 should be 1, got %d", im.GetIdx(0))
	}
	if im.GetIdx(3) != 4 {
		t.Errorf("Index 3 should be 4, got %d", im.GetIdx(3))
	}
	if im.GetKey(42) != 4 {
		t.Errorf("Key 42 should be 4, got %d", im.GetKey(42))
	}
	if im.GetKey(0) != 0 {
		t.Errorf("Key 0 should be 0, got %d", im.GetKey(0))
	}

	im.RemoveIdx(2)
	fmt.Println(im)
	if im.Len() != 4 {
		t.Errorf("Expected len 4, got %d", im.Len())
	}
	if im.GetIdx(0) != 1 {
		t.Errorf("Index 0 should be 1, got %d", im.GetIdx(0))
	}
	if im.GetIdx(1) != 2 {
		t.Errorf("Index 1 should be 2, got %d", im.GetIdx(1))
	}
	if im.GetIdx(2) != 4 {
		t.Errorf("Index 2 should be 4, got %d", im.GetIdx(2))
	}
	if im.GetKey(42) != 4 {
		t.Errorf("Key 42 should be 4, got %d", im.GetKey(42))
	}

	im.RemoveKey(42)
	fmt.Println(im)
	if im.Len() != 3 {
		t.Errorf("Expected len 3, got %d", im.Len())
	}
	if im.GetIdx(0) != 1 {
		t.Errorf("Index 0 should be 1, got %d", im.GetIdx(0))
	}
	if im.GetIdx(1) != 2 {
		t.Errorf("Index 1 should be 2, got %d", im.GetIdx(1))
	}
	if im.GetIdx(2) != 5 {
		t.Errorf("Index 2 should be 5, got %d", im.GetIdx(2))
	}
	if im.GetKey(42) != 0 {
		t.Errorf("Key 42 should be 0, got %d", im.GetKey(42))
	}

	im.RemoveKey(42)
	if im.Len() != 3 {
		t.Errorf("Expected len 3, got %d", im.Len())
	}

	im.RemoveIdx(0)
	fmt.Println(im)
	if im.Len() != 2 {
		t.Errorf("Expected len 2, got %d", im.Len())
	}
	if im.GetIdx(0) != 2 {
		t.Errorf("Index 0 should be 2, got %d", im.GetIdx(0))
	}
	if im.GetIdx(1) != 5 {
		t.Errorf("Index 1 should be 5, got %d", im.GetIdx(1))
	}
	if im.GetKey(42) != 0 {
		t.Errorf("Key 42 should be 0, got %d", im.GetKey(42))
	}

	im.RemoveIdx(1)
	fmt.Println(im)
	if im.Len() != 1 {
		t.Errorf("Expected len 1, got %d", im.Len())
	}
	if im.GetIdx(0) != 2 {
		t.Errorf("Index 0 should be 2, got %d", im.GetIdx(0))
	}

	im.RemoveIdx(0)
	fmt.Println(im)
	if im.Len() != 0 {
		t.Errorf("Expected len 0, got %d", im.Len())
	}

	im2 := New[string, float64]()
	im2.InsertKey("pi", 3.14159)
	im2.InsertKey("e", 2.71828)
	im2.Append(42)

	fmt.Println(im2)

	if im2.Len() != 3 {
		t.Errorf("Expected len 3, got %d", im2.Len())
	}
	if im2.GetIdx(0) != 3.14159 {
		t.Errorf("Index 0 should be 3.14159, got %f", im2.GetIdx(0))
	}
	if im2.GetIdx(1) != 2.71828 {
		t.Errorf("Index 1 should be 2.71828, got %f", im2.GetIdx(1))
	}
	if im2.GetIdx(2) != 42 {
		t.Errorf("Index 2 should be 42, got %f", im2.GetIdx(2))
	}
	if im2.GetKey("pi") != 3.14159 {
		t.Errorf("Key pi should be 3.14159, got %f", im2.GetKey("pi"))
	}
	if im2.GetKey("e") != 2.71828 {
		t.Errorf("Key e should be 2.71828, got %f", im2.GetKey("e"))
	}
	if im2.GetKey("42") != 0 {
		t.Errorf("Key 42 should be 0, got %f", im2.GetKey("42"))
	}
	if im2.GetIdxKey(0) != "pi" {
		t.Errorf("Index 0 should be pi, got %s", im2.GetIdxKey(0))
	}
	if im2.GetIdxKey(1) != "e" {
		t.Errorf("Index 1 should be e, got %s", im2.GetIdxKey(1))
	}
	if im2.GetIdxKey(2) != "" {
		t.Errorf("Index 2 should be \"\", got %s", im2.GetIdxKey(2))
	}
}
