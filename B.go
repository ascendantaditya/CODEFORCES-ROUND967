// Generate Permutation
package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(n int) {
	if n%2 == 1 {
		r := n
		l := 1
		for i := 1; i <= n; i++ {
			if i%2 == 1 {
				fmt.Printf("%d ", r)
				r--
			} else {
				fmt.Printf("%d ", l)
				l++
			}
		}
		fmt.Println()
	} else {
		fmt.Println(-1)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for t > 0 {
		var n int
		fmt.Fscan(reader, &n)
		solve(n)
		t--
	}
}
