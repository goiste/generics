## Generics

some generics-based helper functions for working with slices and sets (go 1.18+)

Usage:
```go
package main

import (
	"fmt"
	"generics/sets"
	"generics/slices"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Reverse(x)) // [5 4 3 2 1]
	f := slices.Convert[int, float64](x)
	fmt.Printf("%T\n", f[0])   // float64
	fmt.Println(slices.Min(f)) // 1
	fmt.Println(slices.Max(f)) // 5
	fmt.Println(slices.Sum(f)) // 15
	
	f = []float64{0.000001, 0.02, 0.300000003}
	formatted := slices.Format(f, "%.2f")
	fmt.Println(formatted) // [0.00 0.02 0.30]

	genInt := slices.SequenceGenerator(10, -1)
	fmt.Println(genInt()) // 10
	fmt.Println(genInt()) // 9
	fmt.Println(genInt()) // 8
	// ...

	s := sets.Make[string]("one", "two", "three")
	s.Add("three", "four", "five")
	s.Diff(sets.Make[string]("four", "five"))
	s.Map(func(str string) string {
		return str + "!"
	})
	fmt.Println(s.Has("one!")) // true
	s.Delete("one!")
	fmt.Println(s.Len()) // 2
}
```