package main

import (
	"spider_simple/engine"
	"spider_simple/parser"
)

// TODO：go get github.com/PuerkitoBio/goquery
func main() {
	url := "https://newcar.xcar.com.cn/car/"
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseIndex,
	})
}
