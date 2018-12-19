package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverseList(&a)
	fmt.Println(a)
	rotate(a[:],6)
	fmt.Println(a)
}

func reverseList(l *[6]int) {
	for i, j := 0, len(*l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
}

func rotate(s []int, n int) {
	if len(s) == 0 || n >= len(s) || n <= 0 {
		return
	}
	ss := make([]int,0)
	ss = append(ss, s...)
	for i := 0; i <= len(s)-1; i++ {
		if  i+n <= len(s)-1 {
			s[i] = ss[i+n]
		} else {
			s[i] = ss[n-1-(len(s)-1-i)]
		}
	}
}