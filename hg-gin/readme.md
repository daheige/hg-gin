# gin action mvc
  基于gin封装的mvc框架,新增了logic,service两个层.
# 网站目录
```
    .
    ├── app.go
    ├── application
    │   ├── config
    │   ├── controller
    │   ├── logic
    │   ├── middleware
    │   ├── model
    │   ├── routes
    │   ├── service
    │   └── views
    ├── bin
    │   └── app-init.sh
    ├── common
    │   └── slog
    ├── docs
    │   ├── nginx.conf
    │   └── readme.md
    ├── public
    │   └── readme.md
    ├── readme.md
    ├── runtime
    └── vendor
        ├── github.com
        ├── gopkg.in
        └── vendor.json
 ```
 # nginx反向代理
    1. 配置nginx(mygo.conf)参考nginx.conf
    2. 配置/etc/hosts
        127.0.0.1  www.mygo.com mygo.com *.mygo.com
    3. 执行go run app.go 访问http://mygo.com/
        {"code":200,"data":["php","go"],"message":"welcome hg-gin page"}
    4. 查看access.log
        heige@daheige:/web/wwwlogs$ tailf  mygo.com-access.log
# 压力测试wrk
    $ wrk -c 600000 -t 2 -d 5 --latency http://mygo.com/v1/hg
    Running 5s test @ http://mygo.com/v1/hg
      2 threads and 600000 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency   231.98ms  480.14ms   1.59s    89.30%
        Req/Sec     2.40k     1.98k    5.17k    44.44%
      Latency Distribution
         50%   24.95ms
         75%   42.99ms
         90%    1.56s 
         99%    1.58s 
      2357 requests in 5.16s, 651.40KB read
      Socket errors: connect 598981, read 0, write 0, timeout 161
    Requests/sec:    456.80
    Transfer/sec:    126.24KB
    
    1. 空跑
    wrk -c 100 -d 10s -t8 http://localhost:8080
    Running 10s test @ http://localhost:8080
      8 threads and 100 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency     4.59ms    4.57ms  75.50ms   87.32%
        Req/Sec     3.12k   454.37     5.01k    73.38%
      248620 requests in 10.05s, 44.34MB read
    Requests/sec:  24736.41
    Transfer/sec:      4.41MB
    100个连接请求10s,qps 24736

    2. 从redis中获取数据
    wrk -c 100 -d 10s -t8 http://localhost:8080/v1/get-user
    Running 10s test @ http://localhost:8080/v1/get-user
      8 threads and 100 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency    61.40ms   64.90ms 474.54ms   84.05%
        Req/Sec   257.15    323.00     1.06k    80.98%
      20545 requests in 10.06s, 3.25MB read
    Requests/sec:   2041.45
    Transfer/sec:    330.94KB
    100个连接,请求10s qps 2041

# 关于wrk
    wrk，简单易用，没有Load Runner那么复杂，他和 apache benchmark（ab）同属于性能测试工具，但是比 ab 功能更加强大，并且可以支持lua脚本来创建复杂的测试场景。

    wrk 的一个很好的特性就是能用很少的线程压出很大的并发量， 原因是它使用了一些操作系统特定的高性能 I/O 机制, 比如 select, epoll, kqueue 等。 其实它是复用了 redis 的 ae 异步事件驱动框架. 确切的说 ae 事件驱动框架并不是 redis 发明的, 它来至于 Tcl的解释器 jim, 这个小巧高效的框架, 因为被 redis 采用而更多的被大家所熟知.

    wrk GitHub 源码：https://github.com/wg/wrk
    ubuntu安装 apt-get install wrk
    wrk里面各个参数什么意思？

    -t 需要模拟的线程数
    -c 需要模拟的连接数
    --timeout 超时的时间
    -d 测试的持续时间
    结果：

    Latency：响应时间
    Req/Sec：每个线程每秒钟的完成的请求数

    Avg：平均
    Max：最大
    Stdev：标准差
    +/- Stdev： 正负一个标准差占比

    标准差如果太大说明样本本身离散程度比较高. 有可能系统性能波动很大.

    如果想看响应时间的分布情况可以加上--latency参数

    我们的模拟测试的时候需要注意，一般线程数不宜过多，核数的2到4倍足够了。 
    多了反而因为线程切换过多造成效率降低， 因为 wrk 
    不是使用每个连接一个线程的模型， 而是通过异步网络 I/O 提升并发量。 
    所以网络通信不会阻塞线程执行，这也是 wrk 
    可以用很少的线程模拟大量网路连接的原因。
    在 wrk 的测试结果中，有一项为Requests/sec，我们一般称之为QPS（每秒请求数）
    ，这是一项压力测试的性能指标，通过这个参数我们可以看出应用程序的吞吐量。
