package main

import (
	"fmt"
	"sort"
)

// 自定义类型
type myIntSlice []int

// 将空值 nil 转换为 *myIntSlice 类型，再转换为 sort.Interface 接口，如果转换失败，说明 myIntSlice 并没有实现 sort.Interface 接口的所有方法。
var _ sort.Interface = (*myIntSlice)(nil)

func (s myIntSlice) Len() int           { return len(s) }
func (s myIntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s myIntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	// []int, []float, []string 三种基础类型的切片排序, 默认是升序排序
	intSlice := make([]int, 0)
	intSlice = append(intSlice, 3, 1, 6, 8, 9)
	sort.Ints(intSlice)
	fmt.Printf("sort intslice %+v\n", intSlice)

	floatSlice := make([]float64, 0)
	floatSlice = append(floatSlice, 1.3, 6.3, 0.1, 3.4, 5.1)
	sort.Float64s(floatSlice)
	fmt.Printf("sort floatslice %+v\n", floatSlice)

	stringSlice := make([]string, 0)
	stringSlice = append(stringSlice, "ak", "cc", "ab", "ee", "ca")
	sort.Strings(stringSlice)
	fmt.Printf("sort stringSlice %+v\n", stringSlice)

	// 使用通用排序函数, 需使用自定义类型, 实现sortLen,Less,Swap三个函数
	intSlice2 := make(myIntSlice, 0)
	intSlice2 = append(intSlice2, 3, 1, 6, 8, 9)
	sort.Sort(intSlice2)
	fmt.Printf("sort common intslice2 %+v\n", intSlice2)

	// 改为降序排序, 需使用自定义类型, 实现Len,Less,Swap三个函数
	intSlice3 := make(myIntSlice, 0)
	intSlice3 = append(intSlice3, 3, 1, 6, 8, 9)
	sort.Sort(sort.Reverse(intSlice3))
	fmt.Printf("sort reverse intslice2 %+v\n", intSlice3)
}
