package engine

import (
	"fmt"
	"spider_simple/fetcher"
)

func Run(seeds ...Request) {
	// 定义request切片保存所有请求
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 遍历所有请求，根据对应的解析器进行解析
	emptyDetail := Details{}
	for len(requests) > 0 {
		// 拿第一个元素并将其从切片中删除
		r := requests[0]
		requests = requests[1:]

		//in := make(chan []Request)
		//out := make(chan []ParserResult)
		parserResult := worker(r)

		for _, item := range parserResult.Items {
			if item == emptyDetail { // 没内容不显示
				continue
			}
			fmt.Printf("Parser Item: %v\n", item)
		}

		// 把获取到的url加到requests中继续走流程
		requests = append(requests, parserResult.Requests...)
	}
}

func worker(r Request) ParserResult {
	// 获取页面
	content, err := fetcher.Fetcher(r.Url)
	if err != nil {
		fmt.Printf("Fetching error: %v\n", err)
	}

	// 根据对应的解析器解析页面内容
	parserResult := r.ParserFunc(content)
	if err != nil {
		fmt.Printf("Parse Error: %v\n", err)
	}
	return parserResult
}
