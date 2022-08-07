package main

import (
	"fmt"
	"sort"
)

// slice-sort-in-go/sort_int_slice.go
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type dto struct {
	A int
	B int
}

func main() {
	src := []int{1, 3, 5, 7, 9}
	dest := make([]int, len(src)) //无容量则无法进行深复制
	copy(dest, src)
	fmt.Println(dest) // [1 3 5 7 9]

	dest2 := make([]int, len(src)-2)
	copy(dest2, src)
	fmt.Println(dest2) // [1 3 5]

	sl := IntSlice([]int{89, 14, 8, 9, 17, 56, 95, 3})
	fmt.Println(sl) // [89 14 8 9 17 56 95 3]
	sort.Sort(sl)
	fmt.Println(sl) // [3 8 9 14 17 56 89 95]

	sort.Stable(sl)
	fmt.Println(sl) // [3 8 9 14 17 56 89 95]

	sort.Sort(sort.Reverse(sl)) // 降序排序
	fmt.Println(sl)             // [95 89 56 17 14 9 8 3]

	// 结构体自定义排序
	dtoSl := make([]dto, 0)
	dtoSl = append(dtoSl, dto{A: 3, B: 4})
	dtoSl = append(dtoSl, dto{A: 6, B: 2})
	dtoSl = append(dtoSl, dto{A: 1, B: 5})
	dtoSl = append(dtoSl, dto{A: 3, B: 2})

	sort.Slice(dtoSl, func(i, j int) bool {
		if dtoSl[i].A != dtoSl[j].A {
			return dtoSl[i].A < dtoSl[j].A
		}

		return dtoSl[i].B > dtoSl[j].B
	})

	fmt.Printf("%+v\n", dtoSl)
}
