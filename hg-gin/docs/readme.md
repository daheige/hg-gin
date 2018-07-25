# golang+nginx反向代理
    1. 配置nginx(mygo.conf)参考nginx.conf
    2. 配置/etc/hosts
        127.0.0.1  www.mygo.com mygo.com *.mygo.com
    3. 执行go run app.go 访问http://mygo.com/
        {"code":200,"data":["php","go"],"message":"welcome hg-gin page"}
    4. 查看access.log
        heige@daheige:/web/wwwlogs$ tailf  mygo.com-access.log 
         127.0.0.1 - - [08/May/2018:22:38:51 +0800] "GET / HTTP/1.1" 200 64 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36"

# pprof
    go get -u github.com/DeanThompson/ginpprof

    然后在app.go中加入
    ginpprof.Wrapper(router)

    通过浏览器访问http://localhost:8080/debug/pprof/
    /debug/pprof/
    profiles:
    0   block
    4   goroutine
    3   heap
    0   mutex
    7   threadcreate

    full goroutine stack dump

# net/http/pprof
    net/http/pprof
    如果程序为 httpserver 类型， 则只需要导入该包:

    import _ "net/http/pprof"
     

    如果 httpserver 使用 go-gin 包，而不是使用默认的 http 包启动，则需要手动添加 /debug/pprof 对应的 handler，github 有封装好的模版:

    import "github.com/DeanThompson/ginpprof"
    ...
    router := gin.Default()
    ginpprof.Wrap(router)
    ...
    导入包重新编译程序后运行，在浏览器中访问 http://host:port/debug/pprof 可以看到如下信息
    /debug/pprof/

    profiles:
    0    block
    62    goroutine
    427    heap
    0    mutex
    12    threadcreate

    full goroutine stack dump

# ginpprof查看heap信息
    需要提前安装好dot安装命令:sudo apt-get install graphviz
    go tool pprof http://127.0.0.1:8080/debug/pprof/heap
    (pprof) web #生成web界面图片
    (pprof) top10 #查看前十条

    pprof mem 分析
    同时 pprof 也支持内存相关数据分析

    --inuse_space 分析常驻内存
    复制代码
    go tool pprof -alloc_space http://127.0.0.1:8080/debug/pprof/heap
    (pprof) top
    Showing nodes accounting for 8055.94kB, 100% of 8055.94kB total
    Showing top 10 nodes out of 24
          flat  flat%   sum%        cum   cum%
     2707.76kB 33.61% 33.61%  5175.02kB 64.24%  compress/flate.NewWriter /usr/local/go/src/compress/flate/deflate.go
     2368.55kB 29.40% 63.01%  2368.55kB 29.40%  runtime/pprof.writeGoroutineStacks /usr/local/go/src/runtime/pprof/pprof.go
     1301.24kB 16.15% 79.17%  1301.24kB 16.15%  compress/flate.(*compressor).init /usr/local/go/src/compress/flate/deflate.go
     1166.01kB 14.47% 93.64%  1166.01kB 14.47%  compress/flate.(*compressor).init /usr/local/go/src/compress/flate/deflatefast.go
     (pprof) top 20s
     Focus expression matched no samples
     Showing nodes accounting for 0, 0% of 1.16MB total
      flat  flat%   sum%        cum   cum%


# go-torch
    uber 开源的工具 go-torch，能让我们将 profile 信息转换成火焰图
    go get github.com/uber/go-torch
    安装好 go-torch 后，运行

    go-torch -u http://127.0.0.1:8080
    生成 CPU 火焰图 https://www.cnblogs.com/upyun/p/8526925.html

# runtime/pprof用法(mytest/demo.go)
    package main

    import (
        "flag"
        "fmt"
        "log"
        "os"
        "runtime"
        "runtime/pprof"
        "sync"
    )

    var cpuprofile string
    var memprofile string

    func init() {
        flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile `file`")
        flag.StringVar(&memprofile, "memprofile", "", "write memory profile to `file`")
        flag.Parse()
    }

    func main() {
        if cpuprofile != "" {
            f, err := os.Create(cpuprofile)
            if err != nil {
                log.Fatal("could not create cpuprofile")
            }

            if err := pprof.StartCPUProfile(f); err != nil {
                log.Fatal("could not start cpu pprofile:", err)
            }

            defer pprof.StopCPUProfile()
        }

        if memprofile != "" {
            f, err := os.Create(memprofile)
            if err != nil {
                log.Fatal("could not create memory profile: ", err)
            }

            runtime.GC()

            if err := pprof.WriteHeapProfile(f); err != nil {
                log.Fatal("could not write memory profile: ", err)
            }
            f.Close()
        }

        var wg sync.WaitGroup
        var i = 0
        for {
            runtime.Gosched()
            if i > 1000000 {
                break
            }

            wg.Add(1)
            go func(i int) {
                defer wg.Done()
                fmt.Println("当前值", i)
            }(i)
            i++
        }

        wg.Wait()
        fmt.Println("finished")
    }

    /**
    运行程序

    go run demo.go -cpuprofile cpu.prof -memprofile mem.prof
    可以得到 cpu.prof 和 mem.prof 文件，使用 go tool pprof 分析。
    go tool pprof logger cpu.prof
    go tool pprof logger mem.prof
    $ go tool pprof demo mem.prof
    demo: open demo: no such file or directory
    fetched 1 profiles out of 2
    Local symbolization failed for demo: open /tmp/go-build365535023/command-line-arguments/_obj/exe/demo: no such file or directory
    Some binary filenames not available. Symbolization may be incomplete.
    Try setting PPROF_BINARY_PATH to the search path for local binaries.
    File: demo
    Type: inuse_space
    Time: Jul 25, 2018 at 10:17pm (CST)
    Entering interactive mode (type "help" for commands, "o" for options)
    (pprof) top 10

    内存分析
    $ go tool pprof demo cpu.prof
    demo: open demo: no such file or directory
    fetched 1 profiles out of 2
    Local symbolization failed for demo: open /tmp/go-build365535023/command-line-arguments/_obj/exe/demo: no such file or directory
    Some binary filenames not available. Symbolization may be incomplete.
    Try setting PPROF_BINARY_PATH to the search path for local binaries.
    File: demo
    Type: cpu
    Time: Jul 25, 2018 at 10:17pm (CST)
    Duration: 8.96s, Total samples = 15.90s (177.41%)
    Entering interactive mode (type "help" for commands, "o" for options)
    (pprof) top 10
    Showing nodes accounting for 8010ms, 50.38% of 15900ms total
    Dropped 148 nodes (cum <= 79.50ms)
    Showing top 10 nodes out of 125
          flat  flat%   sum%        cum   cum%
        2490ms 15.66% 15.66%     2630ms 16.54%  syscall.Syscall /usr/local/go/src/syscall/asm_linux_amd64.s
        1060ms  6.67% 22.33%     4770ms 30.00%  runtime.gentraceback /usr/local/go/src/runtime/traceback.go
        1040ms  6.54% 28.87%     1040ms  6.54%  runtime.futex /usr/local/go/src/runtime/sys_linux_amd64.s
         960ms  6.04% 34.91%     2110ms 13.27%  runtime.pcvalue /usr/local/go/src/runtime/symtab.go
         920ms  5.79% 40.69%     1110ms  6.98%  runtime.step /usr/local/go/src/runtime/symtab.go
         420ms  2.64% 43.33%      480ms  3.02%  runtime.stackpoolalloc /usr/local/go/src/runtime/stack.go
         370ms  2.33% 45.66%      640ms  4.03%  runtime.scanobject /usr/local/go/src/runtime/mgcmark.go
         330ms  2.08% 47.74%      360ms  2.26%  runtime.heapBitsForObject /usr/local/go/src/runtime/mbitmap.go
         220ms  1.38% 49.12%      530ms  3.33%  runtime.scanblock /usr/local/go/src/runtime/mgcmark.go
         200ms  1.26% 50.38%      200ms  1.26%  runtime.findfunc /usr/local/go/src/runtime/symtab.go
    (pprof) top 2s
    Focus expression matched no samples
    Showing nodes accounting for 0, 0% of 15.90s total
          flat  flat%   sum%        cum   cum%

