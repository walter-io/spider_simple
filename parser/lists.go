package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"spider_simple/config"
	"spider_simple/engine"
)

const listsRe = `<a href="/([\w]+)/" target="_blank"[^>]*>([^<]+)</a>`

/**
 * 列表解析器：提取详情页url
 */
func ParseLists(content []byte) engine.ParserResult {
	// 获取dom对象
	result := engine.ParserResult{}
	reader := bytes.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Printf("Goquery list new reader error: %v\n", err)
	}

	// 获取url
	doc.Find(".table_bord").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.
			Find("td").Eq(0).
			Find("p").Eq(0).
			Find("a").Eq(0).
			Attr("href")
		// 这部分逻辑是：给url，并且指定该页面用哪个解析器解析
		result.Items = append(result.Items, engine.Details{})
		result.Requests = append(result.Requests, engine.Request{
			Url:        config.DomainName + url,
			ParserFunc: ParseDetail,
		})
	})
	return result
}
