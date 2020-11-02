#### 简单爬虫：爬取爱卡汽车网汽车列表页(2020-11-2有效)
##### 爬取连接：https://newcar.xcar.com.cn/car
##### 本爬虫有一个并发版，架构大致相同，请参考：https://github.com/sgyouyou/spider_goroutine

##### 安装库
````
> go mod init
> go get github.com/PuerkitoBio/goquery
> go get golang.org/x/net/html/charset
> go get golang.org/x/text/encoding
 ````
##### 目录结构
````
> config            # 配置目录
> engine            # 中控引擎
> fetcher           # 抓取器
> parser            # 解析器
> main.go           # 入口文件
````

##### 总体的逻辑是：
````
> main.go中调用中控引擎
> 引擎中调用抓取器抓取页面内容，并提交给解析器解析出我们想要的内容并返回给引擎
````
 
##### 直接运行main.go即可在终端看到爬起结果

