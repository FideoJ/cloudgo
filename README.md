## Description

《服务计算》作业四：用 Go 开发类似 [cloudgo](http://blog.csdn.net/pmlpml/article/details/78404838) 的 web 服务程序

## Setup

```
go get github.com/FideoJ/cloudgo
$GOPATH/bin/cloudgo -p 8080
```

## 关于框架选择

[gin](https://gin-gonic.github.io/gin/)  


选用合适的框架，能让web应用开发者将主要精力放在业务开发上。  
gin 为各种常见的web应用基础需求提供了成熟的解决方案，如日志、路由等。

## curl 测试

```
$ curl -v localhost:8080/hello/fideo
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /hello/fideo HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Tue, 23 Jan 2018 01:11:39 GMT
< Content-Length: 34
< 
* Connection #0 to host localhost left intact
{"data":"hi fideo!","status":"OK"}
```


```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /hello/:name              --> main.createServer.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2018/01/23 - 09:11:26 | 200 |      88.214µs |       127.0.0.1 | GET      /hello/fideo
```

## ab 测试结果

### 测试命令

```
$ ab -n 1000 -c 100 http://localhost:8080/hello/fideo
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /hello/fideo
Document Length:        34 bytes

Concurrency Level:      100
Time taken for tests:   0.173 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      157000 bytes
HTML transferred:       34000 bytes
Requests per second:    5784.56 [#/sec] (mean)
Time per request:       17.287 [ms] (mean)
Time per request:       0.173 [ms] (mean, across all concurrent requests)
Transfer rate:          886.89 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   2.2      2      13
Processing:     1   14  10.7     10      55
Waiting:        0   12   9.9      8      55
Total:          1   17  10.4     13      57

Percentage of the requests served within a certain time (ms)
  50%     13
  66%     17
  75%     20
  80%     23
  90%     34
  95%     40
  98%     46
  99%     48
 100%     57 (longest request)
```

- Server Software:        服务器使用的软件，即响应头的 Server 字段
- Server Hostname:        服务器主机名，即请求头的 Host 字段
- Server Port:            服务器端口

- Document Path:          文档路径，即请求头的请求路径
- Document Length:        文档长度，即响应头的 Content-Length 字段

- Concurrency Level:      并发等级，即并发数
- Time taken for tests:   测试花费的时间
- Complete requests:      完成的请求数
- Failed requests:        失败的请求数
- Total transferred:      传输的字节数
- HTML transferred:       传输的 HTML 报文体的字节数
- Requests per second:    平均每秒的请求数
- Time per request:       平均每个请求花费的时间
- Time per request:       考虑并发，平均每个请求花费的时间
- Transfer rate:          平均每秒传输的千字节数

- Connection Times (ms) 传输时间统计

```
            min最小  mean期望[+/-sd]标准差 median中位数   max最大
Connect:    连接时间
Processing: 处理时间
Waiting:    等待时间
Total:      总时间
```

- Percentage of the requests served within a certain time (ms) 一定时间内服务了的请求数所占的百分比

```
  50%     N    N 毫秒内服务了50%的请求
  ...
 100%     M    M 毫秒内服务了100%的请求(最长请求)
```