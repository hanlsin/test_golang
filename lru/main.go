package main

import "fmt"

const LRU_MAX = 5

type LRU struct {
	m map[int]string
	l []*int
}

var lru *LRU

func addLRU(key int, value string) {
	idx := len(lru.l) - 1
	_, ok := lru.m[key]
	if idx < 0 || !ok {
		lru.m[key] = value
		lru.l = append(lru.l, &key)
	} else {
		// if already last, return
		if key == *lru.l[idx] {
			fmt.Println("Already it's last")
			return
		}

		for i, v := range lru.l {
			if key == *v {
				fmt.Println(i)
				lru.l = append(lru.l[:i], lru.l[i+1:]...)
				lru.l = append(lru.l, &key)
				break
			}
		}
	}

	// check max number
	if len(lru.l) > LRU_MAX {
		fmt.Println("over max")
		delete(lru.m, *lru.l[0])
		lru.l = lru.l[1:]
	}
}

func printLRU() {
	fmt.Printf("LRU = ")
	for _, v := range lru.l {
		fmt.Printf("[%d] %s, ", *v, lru.m[*v])
	}
	fmt.Println()
	fmt.Println()
}

func main() {
	lru = &LRU{
		m: make(map[int]string),
		l: make([]*int, 0, LRU_MAX+1),
	}

	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := 3
	fmt.Println(list[:i-1])
	fmt.Println(list[i:])

	addLRU(10, "10")
	addLRU(11, "11")
	addLRU(12, "12")
	addLRU(13, "13")
	addLRU(13, "13")
	printLRU()

	addLRU(11, "11")
	addLRU(11, "11")
	printLRU()

	addLRU(14, "14")
	addLRU(15, "15")
	printLRU()

	addLRU(10, "10")
	printLRU()

	addLRU(14, "14")
	printLRU()
}
