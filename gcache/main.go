package main

import (
	"fmt"
	"time"

	cache "go_demo/gcache/gcache"

	gcache "github.com/patrickmn/go-cache"
)

type Data struct {
	A int
	B string
	C float32
}

func main() {
	cache.Init()

	cache.Instance.Set("k1", "v1", 0)
	v, exist := cache.Instance.Get("k1")
	if !exist {
		fmt.Println("k1 expired")
	} else {
		fmt.Printf("k1 val: %s\n", v.(string))
	}

	cache.Instance.Set("k2", "v2", 5*time.Second)
	time.Sleep(2 * time.Second)
	v2, exist2 := cache.Instance.Get("k2")
	if !exist2 {
		fmt.Println("k2 expired")
	} else {
		fmt.Printf("k2 val: %s\n", v2.(string))
	}

	time.Sleep(5 * time.Second)
	v2, exist2 = cache.Instance.Get("k2")
	if !exist2 {
		fmt.Println("k2 expired")
	} else {
		fmt.Printf("k2 val: %s\n", v2.(string))
	}

	cache.Instance.Set("k2", "v2", gcache.NoExpiration) // 不过期

	d := Data{A: 11, B: "123", C: 3.12}
	cache.Instance.Set("k3", d, gcache.NoExpiration)
	d2, _ := cache.Instance.Get("k3")
	fmt.Printf("data: %+v\n", d2.(Data))
}
