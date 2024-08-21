package main

import (
    "fmt"
)

const MAXN = 1000

func q(a, b int) int {
    fmt.Printf("? %d %d\n", a, b)
    var r int
    fmt.Scan(&r)
    return r
}

func function(a, b int, vis []int, edges *[][2]int) {
    x := q(a, b)
    if x == a || x == b {
        *edges = append(*edges, [2]int{a, b})
        vis[x] = 1
        vis[b] = 1
        return
    }
    if vis[a] == 1 && vis[x] == 1 {
        function(x, b, vis, edges)
    } else {
        function(a, x, vis, edges)
        function(x, b, vis, edges)
    }
}

func main() {
    var t int
    fmt.Scan(&t)
    for t > 0 {
        t--
        var n int
        fmt.Scan(&n)

        if n == 2 {
            fmt.Println("! 1 2")
            continue
        }

        vis := make([]int, n+1)
        edges := make([][2]int, 0)

        for i := 1; i < n; i++ {
            for j := i + 1; j <= n; j++ {
                if vis[j] == 1 {
                    continue
                }
                function(i, j, vis, &edges)
            }
        }
        fmt.Print("!")
        for _, edge := range edges {
            fmt.Printf(" %d %d", edge[0], edge[1])
        }
        fmt.Println()
    }
}
