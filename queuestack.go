package main

import (
   "fmt"
)

func main() {
   q := make([]int,0)
   // 3 pushes
   q = append(q, 1)
   q = append(q, 2)
   q = append(q, 3)
   fmt.Println("push 1, 2, 3, to queue:",q)
   // 2 pops
   var v int
   v = q[len(q)-1]
   q = q[0:len(q)-1]
   fmt.Println("pop from queue, ", v)
   v = q[len(q)-1]
   q = q[0:len(q)-1]
   fmt.Println("pop from queue, ",v)
   fmt.Println("queue: ", q)
}
