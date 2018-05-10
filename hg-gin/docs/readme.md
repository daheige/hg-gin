# golang+nginx反向代理
    1. 配置nginx(mygo.conf)参考nginx.conf
    2. 配置/etc/hosts
        127.0.0.1  www.mygo.com mygo.com *.mygo.com
    3. 执行go run app.go 访问http://mygo.com/
        {"code":200,"data":["php","go"],"message":"welcome hg-gin page"}
    4. 查看access.log
        heige@daheige:/web/wwwlogs$ tailf  mygo.com-access.log 
         127.0.0.1 - - [08/May/2018:22:38:51 +0800] "GET / HTTP/1.1" 200 64 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36"


