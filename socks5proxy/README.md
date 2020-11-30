# socks5 proxy

## test result

1. start mock http server `go run httpserver.go`
1. start socks5 proxy server `go run socks5proxy.go`
1. bench
    - `gobench -u http://127.0.0.1:8080/ping -c 100 -n 1000000`
    - `gobench -u http://127.0.0.1:8080/ping -c 100 -n 1000000 -proxy="socks5://127.0.0.1:1080"`
    - 本机测试结果，性能下降一半，不是[150行Go实现高性能socks5代理](https://mp.weixin.qq.com/s/WjRRCU3xKvDRKgru9dZ7hg)博主说的下降一点点

result

```bash
❯ gobench -u http://127.0.0.1:8080/ping -c 100 -n 1000000
Dispatching 100 goroutines at 2020-11-30 14:49:22.262
Waiting for results...

Total Requests:                 1000000 hits
Successful requests:            1000000 hits
Network failed:                 0 hits
Bad requests(!2xx):             0 hits
Successful requests rate:       119862 hits/sec
Read throughput:                14 MiB/sec
Write throughput:               10 MiB/sec
Test time:                      8.343s(2020-11-30 14:49:22.262-14:49:30.605)

❯ gobench -u http://127.0.0.1:8080/ping -c 100 -n 1000000 -proxy="socks5://127.0.0.1:1080"
Dispatching 100 goroutines at 2020-11-30 14:50:21.182
Waiting for results...

Total Requests:                 1000000 hits
Successful requests:            1000000 hits
Network failed:                 0 hits
Bad requests(!2xx):             0 hits
Successful requests rate:       52816 hits/sec
Read throughput:                6.0 MiB/sec
Write throughput:               4.5 MiB/sec
Test time:                      18.934s(2020-11-30 14:50:21.182-14:50:40.116)

```

## thanks

1. [150行Go实现高性能socks5代理](https://mp.weixin.qq.com/s/WjRRCU3xKvDRKgru9dZ7hg)
1. [felix021/socks5_proxy.go](https://gist.github.com/felix021/7f9d05fa1fd9f8f62cbce9edbdb19253)
1. [https://github.com/cnlh/benchmark](https://github.com/cnlh/benchmark)