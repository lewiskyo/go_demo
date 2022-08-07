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

func worker(group *sync.WaitGroup, tasks chan string, results chan string) {
	for task := range tasks {
		if task == "" {
			close(tasks)
		} else {
			respBody, err := request(task)
			if err != nil {
				fmt.Printf("error occured in request: %s\n", task)
				results <- err.Error()
			} else {
				results <- string(respBody)
			}
		}
	}
	group.Done()
}

const routineLimit = 2 // 限制只能开启两个协程处理

func main() {
	var numberTasks = [5]string{"1345675544", "131938575", "1341931788", "1343434343", "123423242"}
	client = &http.Client{}

	beg := time.Now()
	wg := &sync.WaitGroup{}
	tasks := make(chan string)   // 接收任务用通信chan
	results := make(chan string) // 接口请求结果用chan

	go func() {
		for result := range results {
			if result == "" {
				close(results)
			} else {
				fmt.Println("result: ", result)
			}
		}
	}()

	// 开启指定数量工作协程
	for i := 0; i < routineLimit; i++ {
		wg.Add(1)
		go worker(wg, tasks, results)
	}

	// 分发任务
	for _, task := range numberTasks {
		tasks <- task
	}

	tasks <- "" // worker结束标志
	wg.Wait()   // 同步等待
	results <- ""

	fmt.Printf("time consumed: %fs\n", time.Now().Sub(beg).Seconds())
}
