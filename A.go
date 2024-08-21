// Make all Equal
package main

import (
	"fmt"
)

const N = 105

func solve() {
	var n int
	fmt.Scan(&n)

	cnt := make([]int, N)
	ans := 0

	for i := 1; i <= n; i++ {
		var x int
		fmt.Scan(&x)
		cnt[x]++
		if cnt[x] > ans {
			ans = cnt[x]
		}
	}

	fmt.Println(n - ans)
}

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		solve()
		t--
	}
}
