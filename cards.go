package main

import (
	"fmt"
	"os"
	"sort"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	var n, l, r, num int
	var A, B []int

	fmt.Fscan(os.Stdin, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &num)
		A = append(A, num)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &num)
		B = append(B, num)
	}

	for i := range A {
		if A[i] != B[i] {
			l = i
			break
		}
	}
	for i := range A {
		if A[n-i-1] != B[n-i-1] {
			r = n - i
			break
		}
	}

	sort.Ints(A[l:r])

	if Equal(A, B) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
