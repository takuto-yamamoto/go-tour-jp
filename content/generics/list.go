// +build OMIT

package main

// List は、任意の型の値を持つ単方向連結リストを表現しています
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}