package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var client *http.Client

func request(keyword string) (body []byte, err error) {
	url := fmt.Sprintf("http://127.0.0.1:8080/query?number=%s", keyword)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		data, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("response status code is not OK, response code is %d, body: %s", resp.StatusCode, string(data))
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	var numberTasks = [5]string{"1345675544", "131938575", "1341931788", "1343434343", "123423242"}
	client = &http.Client{}

	beg := time.Now()
	wg := &sync.WaitGroup{}
	// 	在for循环中有多少个任务就go出去多少个协程,没有限制,在查询量固定或者不大的时候没什么问题.
	//  没有用到 channel的特性, 仅利用了多核调度.
	for _, keyword := range numberTasks {
		wg.Add(1)
		go func(keyword string, group *sync.WaitGroup) {
			body, err := request(keyword)
			if err != nil {
				fmt.Printf("error occured in query keyword: %s, error: %s\n", keyword, err.Error())
			} else {
				fmt.Printf("search %s success, data size is %d,\nbody is %s\n", keyword, len(body), string(body))
			}
			group.Done()
		}(keyword, wg)
	}
	wg.Wait()
	fmt.Printf("time consumed: %fs\n", time.Now().Sub(beg).Seconds())
}
