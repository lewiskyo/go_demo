package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

func connectRedis() *redis.Client {
	rd := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	result, err := rd.Ping().Result()
	if err != nil {
		fmt.Printf("ping redis err %s", err)
		return nil
	}
	fmt.Printf("ping redis result: %s\n", result)

	return rd
}

// 测试string占用内存
func testString(client *redis.Client) {
	wg := sync.WaitGroup{}
	beginSlice := []int{}
	begin := 1000000000
	segmentLen := 1500000
	seg := 10
	for i := 0; i < seg; i++ {
		beginSlice = append(beginSlice, begin+segmentLen*i)
	}

	var v interface{}
	v = 1

	for i := 0; i < seg; i++ {
		temp := i
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			beg := beginSlice[idx]
			end := beg + segmentLen - 1
			for j := beg; j <= end; j++ {
				key := strconv.Itoa(j)
				client.Set(key, v, time.Second*86400)
			}
		}(temp)
	}

	wg.Wait()
}

// 测试分段set占用内存
func testSegmentSet(client *redis.Client) {
	wg := sync.WaitGroup{}
	beginSlice := []int{}
	begin := 1000000000
	segmentLen := 150000
	// segmentLen := 100
	seg := 100
	for i := 0; i < seg; i++ {
		beginSlice = append(beginSlice, begin+segmentLen*i)
	}

	for i := 0; i < seg; i++ {
		temp := i
		wg.Add(1)
		go func(idx int) {
			key := fmt.Sprintf("testset%d", idx)
			defer wg.Done()
			beg := beginSlice[idx]
			end := beg + segmentLen - 1
			for j := beg; j <= end; j++ {
				value := strconv.Itoa(j)
				client.SAdd(key, value)
			}
		}(temp)
	}

	wg.Wait()
}

// 测试uuid分段set占用内存
func testUuidSegmentSet(client *redis.Client) {
	wg := sync.WaitGroup{}
	seg := 10000
	per := 17000

	for i := 0; i < seg; i++ {
		wg.Add(1)
		temp := i
		go func(idx int) {
			defer wg.Done()
			key := strconv.Itoa(idx)
			for j := 0; j < per; j++ {
				uuid := uuid.New()
				client.SAdd(key, uuid.String())
			}
		}(temp)
	}

	wg.Wait()
}

// 测试分段bitmap占用内存
func testSegmentBitmap(client *redis.Client) {
	length := 10000000

	for i := 1; i <= 300; i++ {
		key := fmt.Sprintf("testbit%d", i)
		client.SetBit(key, int64(length), 1)
	}
}

func main() {
	rd := connectRedis()
	testUuidSegmentSet(rd)
	// testString(rd)
	// testSegmentSet(rd)
	// testSegmentBitmap(rd)
}
