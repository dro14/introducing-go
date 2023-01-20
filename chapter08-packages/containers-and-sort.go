package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type llist struct {
	list.List
}

func (l llist) Len() int {
	length := 0
	for node := l.Front(); node != nil; node = node.Next() {
		length++
	}
	return length
}

func (l llist) Less(i, j int) bool {
	index := 0
	iNode := l.Front()
	for index != i {
		index++
		iNode = iNode.Next()
	}
	jndex := 0
	jNode := l.Front()
	for jndex != j {
		jndex++
		jNode = jNode.Next()
	}
	return iNode.Value.(int) < jNode.Value.(int)
}

func (l llist) Swap(i, j int) {
	index := 0
	iNode := l.Front()
	for index != i {
		index++
		iNode = iNode.Next()
	}
	jndex := 0
	jNode := l.Front()
	for jndex != j {
		jndex++
		jNode = jNode.Next()
	}
	iNode.Value, jNode.Value = jNode.Value, iNode.Value
}

func printLlist(l *llist) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Print("\n")
}

func main() {
	var n int
	fmt.Print("Enter the size of the list: ")
	fmt.Scan(&n)

	var l llist
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		l.PushBack(rand.Intn(n))
	}

	start := time.Now()
	sort.Sort(llist(l))
	fmt.Println("Time elapsed:", time.Since(start))
}
