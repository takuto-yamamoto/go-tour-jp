// +build OMIT

package main

import "fmt"

// Index は s における x のインデックスを返す、もし見つからない場合は -1 を返す
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v と x は型 T であり、 comparable 制約を持つ
		// なのでここでは == が使用できる
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index は int のスライスに対して動作する
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index は string のスライスに対しても動作する
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}