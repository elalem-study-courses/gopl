package main

import "fmt"

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func rotate(a []int, n int) {
	d := gcd(len(a), n)

	for i := 0; i < d; i++ {
		tmp := a[i]
		j := i
		for {
			k := j + d
			if k >= len(a) {
				k -= len(a)
			}

			if k == i {
				break
			}

			a[j] = a[k]
			j = k
		}

		a[j] = tmp
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	rotate(a, 3)
	fmt.Println(a)
}
