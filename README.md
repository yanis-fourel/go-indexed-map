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
	specialNumbers.InsertKey("pi", 3.14)
	specialNumbers.InsertKey("e", 2.71828)
	specialNumbers.InsertKey("phi", 1.61803)
	specialNumbers.Append(69)
	specialNumbers.InsertKey("sqrt(2)", 1.41421)

	for idx, item := range specialNumbers.Iter() {
		fmt.Printf("index=%d, key='%s', value=%f\n", idx, item.Key, item.Val)
	}

	val := specialNumbers.GetIdx(2)
	fmt.Printf("value at index 2: %f\n", val)

	key := specialNumbers.GetIdxKey(2)
	fmt.Printf("key of value at index 2: %s\n", key)

	val = specialNumbers.GetKey("pi")
	fmt.Printf("value of key 'pi': %f\n", val)

	idx := specialNumbers.GetKeyIdx("pi")
	fmt.Printf("index of key 'pi': %d\n", idx)

	specialNumbers.RemoveIdx(2)
	fmt.Println("removed index 2")
	val = specialNumbers.GetIdx(2)
	fmt.Printf("value at index 2: %f\n", val)

	specialNumbers.RemoveKey("sqrt(2)")
	fmt.Println("removed key 'sqrt(2)'")

	for idx, item := range specialNumbers.Iter() {
		fmt.Printf("index=%d, key='%s', value=%f\n", idx, item.Key, item.Val)
	}

	if specialNumbers.HasKey("pi") {
		fmt.Println("specialNumbers has key 'pi'")
	} else {
		fmt.Println("specialNumbers does not have key 'pi'")
	}
}
```
Running above example produces the following output.
```
index=0, key='', value=42.000000
index=1, key='pi', value=3.140000
index=2, key='e', value=2.718280
index=3, key='phi', value=1.618030
index=4, key='', value=69.000000
index=5, key='sqrt(2)', value=1.414210
value at index 2: 2.718280
key of value at index 2: e
value of key 'pi': 3.140000
index of key 'pi': 1
removed index 2
value at index 2: 1.618030
removed key 'sqrt(2)'
index=0, key='', value=42.000000
index=1, key='pi', value=3.140000
index=2, key='phi', value=1.618030
index=3, key='', value=69.000000
specialNumbers has key 'pi'
```
