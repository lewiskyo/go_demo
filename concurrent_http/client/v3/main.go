package main

import (
	"fmt"
	"go_demo/concurrent_http/client/v3/golimit"
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

const routineLimit = 2 // 限制只能开启两个协程处理

func main() {
	var numberTasks = [5]string{"1345675544", "131938575", "1341931788", "1343434343", "123423242"}
	client = &http.Client{}

	beg := time.Now()
	wg := &sync.WaitGroup{}
	g := golimit.NewG(routineLimit)

	// 巧妙地使用go带缓冲区的通道来实现goroutine控制, 更加简洁, 并且这种方式在多个项目中都可以复用. 不必像v2版本每次都实现一个worker函数.
	for i := 0; i < len(numberTasks); i++ { // 此处不能直接使用for - range形式
		wg.Add(1)
		task := numberTasks[i]
		g.Run(func() {
			respBody, err := request(task)
			if err != nil {
				fmt.Printf("error occured in request: %s\n", task)
			} else {
				fmt.Printf("response data: %s\n", respBody)
			}
			wg.Done()
		})
	}

	wg.Wait() // 同步等待

	fmt.Printf("time consumed: %fs\n", time.Now().Sub(beg).Seconds())
}
