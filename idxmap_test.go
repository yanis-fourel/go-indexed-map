package idxmap

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	im := New()

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
}
