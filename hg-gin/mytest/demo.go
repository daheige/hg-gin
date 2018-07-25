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
*/
