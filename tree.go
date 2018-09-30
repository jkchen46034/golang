package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

/*
    0
   / \
  1   2
 / \   \
3   4   5
 \      /
 7     6

*/

func Equal(a []int, b []int) bool {
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

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	node7 := Node{7, nil, nil}
	node3 := Node{3, nil, &node7}
	node4 := Node{4, nil, nil}
	node6 := Node{6, nil, nil}
	node8 := Node{8, nil, nil}
	node5 := Node{5, &node6, nil}
	node2 := Node{2, &node8, &node5}
	node1 := Node{1, &node3, &node4}
	node0 := Node{0, &node1, &node2}

	fmt.Println(`
        0
       / \
      1   2
      /\  / \
     3  4 8  5
      \     /
      7    6
    `)

	fmt.Print("Infix: ")
	infix(&node0)

	fmt.Print("\nPrefix: ")
	prefix(&node0)

	fmt.Print("\nPostfix: ")
	postfix(&node0)

	fmt.Print("\nBFS: ")
	bfs(&node0)

	fmt.Println("\nHeight of the tree is: ", height(&node0))

	q := make([]int, 0)
	path(&node0, 4, &q)
	fmt.Println("Path from 0 to 4 is ", q)

	q = q[0:0]
	path(&node0, 6, &q)
	fmt.Println("Path from 0 to 6 is ", q)

	q = q[0:0]
	path(&node0, 3, &q)
	fmt.Println("Path from 0 to 3 is ", q)

	q = q[0:0]
	path(&node0, 0, &q)
	fmt.Println("Path from 0 to 0 is ", q)

	fmt.Println("Route from 8 to 5 is  ", route(&node0, 8, 5))
	fmt.Println("Route from 4 to 5 is  ", route(&node0, 4, 5))
	fmt.Println("Route from 7 to 6 is ", route(&node0, 7, 6))
	fmt.Println("Route from 7 to 4 is ", route(&node0, 7, 4))
	fmt.Println("Route from 1 to 8 is ", route(&node0, 1, 8))

	fmt.Println("Longest Expression: ",
		LongestExpression(")((())())))) ()()"))

	allPaths := make([]int, 0)
	fmt.Println("Path to each leaf: ")
	paths(&node0, &allPaths)

	fmt.Println("Another all Path to each leaf: ")
	AllPaths(&node0, make([]int, 0))

	longestPaths := make([][]int, 0)
	LongestPath(&node0, make([]int, 0), &longestPaths)
	for _, path := range longestPaths {
		fmt.Println("Longest path: ", path)
	}

	var dmax int
	height := Diemeter(&node0, &dmax)
	fmt.Println("Diemeter of tree:", dmax, "; height of tree: ", height)

	fmt.Println("LCA of two nodes 3 and 8 is ", LCA(&node0, 3, 8))

	fmt.Println("LCA of two nodes 7 and 4 is ", LCA(&node0, 7, 4))

	fmt.Println("The number of nodes is ", count(&node0))
}

func infix(n *Node) {
	if n == nil {
		return
	}
	infix(n.left)
	visit(n)
	infix(n.right)
}

func prefix(n *Node) {
	if n == nil {
		return
	}
	visit(n)
	prefix(n.left)
	prefix(n.right)
}

func postfix(n *Node) {
	if n == nil {
		return
	}
	postfix(n.left)
	postfix(n.right)
	visit(n)
}

func visit(n *Node) {
	fmt.Print(n.val, " ")
}

func bfs(n *Node) {
	if n == nil {
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, n)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		visit(n)
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
}

func height(n *Node) int {
	if n == nil {
		return 0
	}

	lheight := height(n.left)
	rheight := height(n.right)

	return Max(lheight, rheight) + 1
}

// the number of nodes
func count(n *Node) int {
	if n == nil {
		return 0
	}
	return count(n.left) + count(n.right) + 1
}

func path(n *Node, val int, q *[]int) bool {
	if n == nil {
		return false
	}

	*q = append(*q, n.val)

	if n.val == val {
		return true
	}

	if path(n.left, val, q) || path(n.right, val, q) {
		return true
	}

	*q = (*q)[0 : len(*q)-1]

	return false
}

func route(n *Node, from int, to int) []int {
	q := make([]int, 0)

	list1 := make([]int, 0)
	list2 := make([]int, 0)
	if !path(n, from, &list1) || !path(n, to, &list2) {
		return q
	}

	var i int
	for i = 0; i < len(list1) && i < len(list2); i++ {
		if list1[i] != list2[i] {
			break
		}
	}

	// least common parent
	lca := i - 1

	// write the route to q
	for j := len(list1) - 1; j >= lca; j-- {
		q = append(q, list1[j])
	}

	for j := lca + 1; j < len(list2); j++ {
		q = append(q, list2[j])
	}
	return q
}

func LongestExpression(s string) int {
	stack := make([]rune, 0)
	var maxlen int
	var cnt int
	for _, c := range s {
		switch c {
		case '(':
			stack = append(stack, c)
		case ')':
			pop := stack[len(stack)-1]
			if len(stack) > 0 && pop == '(' {
				cnt = cnt + 1
				stack = stack[0 : len(stack)-1]
			} else {
				if cnt > maxlen {
					maxlen = cnt
				}
				cnt = 0
			}
		}
	}
	return maxlen * 2
}

func paths(n *Node, path *[]int) {
	if n == nil {
		return
	}

	*path = append(*path, n.val)

	if n.left == nil && n.right == nil {
		fmt.Println("path: ", *path)
		return
	}

	if n.left != nil {
		paths(n.left, path)
		*path = (*path)[0 : len(*path)-1]
	}

	if n.right != nil {
		paths(n.right, path)
		*path = (*path)[0 : len(*path)-1]
	}
}

func AllPaths(n *Node, path []int) {
	if n == nil {
		return
	}

	path = append(path, n.val)

	if n.left == nil && n.right == nil {
		fmt.Println("path: ", path)
		return
	}

	AllPaths(n.left, path)
	AllPaths(n.right, path)
}

func appendOrReplace(path []int, longestPaths *[][]int) {
	if len(*longestPaths) == 0 {
		(*longestPaths) = append(*longestPaths, path)
	} else if len(path) > len((*longestPaths)[0]) {
		(*longestPaths)[0] = path
		(*longestPaths) = (*longestPaths)[0:1]
	} else if len(path) == len((*longestPaths)[0]) {
		(*longestPaths) = append(*longestPaths, path)
	}
}

func LongestPath(n *Node, path []int, longestPaths *[][]int) {
	if n == nil {
		return
	}

	path = append(path, n.val)

	if n.left == nil && n.right == nil {
		appendOrReplace(path, longestPaths)
		return
	}

	LongestPath(n.left, path, longestPaths)
	LongestPath(n.right, path, longestPaths)
}

func Diemeter(n *Node, dmax *int) int {
	if n == nil {
		return 0
	}

	lHeight := Diemeter(n.left, dmax)
	rHeight := Diemeter(n.right, dmax)

	height := Max(lHeight, rHeight) + 1
	diemeter := lHeight + rHeight + 1

	*dmax = Max(*dmax, diemeter)

	//fmt.Println(" node ", n.val, " height ", height, "diemeter ", diemeter, "max diemeter " , *dmax)
	return height
}

func LCA(n *Node, from int, to int) int {
	return 0
}

/*
$ go run tree.go

        0
       / \
      1   2
      /\  / \
     3  4 8  5
      \     /
      7    6

Infix: 3 7 1 4 0 8 2 6 5
Prefix: 0 1 3 7 4 2 8 5 6
Postfix: 7 3 4 1 8 6 5 2 0
BFS: 0 1 2 3 4 8 5 7 6
Height of the tree is:  4
Path from 0 to 4 is  [0 1 4]
Path from 0 to 6 is  [0 2 5 6]
Path from 0 to 3 is  [0 1 3]
Path from 0 to 0 is  [0]
Route from 8 to 5 is   [8 2 5]
Route from 4 to 5 is   [4 1 0 2 5]
Route from 7 to 6 is  [7 3 1 0 2 5 6]
Route from 7 to 4 is  [7 3 1 4]
Route from 1 to 8 is  [1 0 2 8]
Longest Expression:  8
Path to each leaf:
path:  [0 1 3 7]
path:  [0 1 4]
path:  [0 2 8]
path:  [0 2 5 6]
Another all Path to each leaf:
path:  [0 1 3 7]
path:  [0 1 4]
path:  [0 2 8]
path:  [0 2 5 6]
Longest path:  [0 1 3 7]
Longest path:  [0 2 5 6]
Diemeter of tree: 7 ; height of tree:  4
LCA of two nodes 3 and 8 is  0
LCA of two nodes 7 and 4 is  0
The number of nodes is  9
*/
