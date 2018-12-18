package main

import "fmt"

func main() {
	fmt.Println(equal([]string{"a", "b"}, []string{"a", "b"}))
}

//slice不支持除 slice == nil 之外的 ==比较
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
