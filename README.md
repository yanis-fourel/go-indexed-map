# go-index-map

Simple, generic implementation of an indexed map in go.  
This data structure is an array where each item has an optional key. If an item
has a key, it is unique throughout the entire array. Elements can be accessed
either through their index or their key.  
See usage for example

# Installation

Install `go-index-map` by executing the following command.
```
go get -v github.com/yanis-fourel/go-indexed-map
```

# Usage

```go
package main

import (
	"fmt"
	"github.com/yanis-fourel/go-indexed-map"
)

func main() {
	specialNumbers := idxmap.New[string, float64]()

	specialNumbers.Append(42)
	specialNumbers.Set("pi", 3.14)
	specialNumbers.Set("e", 2.71828)
	specialNumbers.Set("phi", 1.61803)
	specialNumbers.Append(69)
	specialNumbers.Set("sqrt(2)", 1.41421)

	for idx, item := range specialNumbers.Slice() {
		fmt.Printf("index=%d, key='%s', value=%f\n", idx, item.Key, item.Val)
	}

```
At this point, the output is
```
index=0, key='', value=42.000000
index=1, key='pi', value=3.140000
index=2, key='e', value=2.718280
index=3, key='phi', value=1.618030
index=4, key='', value=69.000000
index=5, key='sqrt(2)', value=1.414210
```
We can access and modify elements
```go
	val := specialNumbers.At(2)        // 2.718280
	key := specialNumbers.GetIdxKey(2) // "e"

	val = specialNumbers.Get("pi")        // 3.14
	idx := specialNumbers.GetKeyIdx("pi") // 1

	has := specialNumbers.HasKey("pi") // true

	length := specialNumbers.Len()           // 6
	lengthKeyed := specialNumbers.LenKeyed() // 4

	// note: Removal preserve order
	specialNumbers.RemoveAt(3)
	val = specialNumbers.At(3)        // 69
	key = specialNumbers.GetIdxKey(3) // ""

	specialNumbers.Remove("pi")

	for idx, item := range specialNumbers.Slice() {
		fmt.Printf("index=%d, key='%s', value=%f\n", idx, item.Key, item.Val)
	}
}
```
Output
```
index=0, key='', value=42.000000
index=1, key='e', value=2.718280
index=2, key='', value=69.000000
index=3, key='sqrt(2)', value=1.414210
```
And we can do any arbitrary operation on the underlying slice like
```go
	s := specialNumbers.Slice()
	sort.Slice(s, func(i, j int) bool { return s[i].Val < s[j].Val })
	specialNumbers = idxmap.From(s)

	for idx, item := range specialNumbers.Slice() {
		fmt.Printf("index=%d, key='%s', value=%f\n", idx, item.Key, item.Val)
	}
```
Output
```
index=0, key='sqrt(2)', value=1.414210
index=1, key='e', value=2.718280
index=2, key='', value=42.000000
index=3, key='', value=69.000000
```
