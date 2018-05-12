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
 ```
    1. 配置nginx(mygo.conf)参考nginx.conf
    2. 配置/etc/hosts
        127.0.0.1  www.mygo.com mygo.com *.mygo.com
    3. 执行go run app.go 访问http://mygo.com/
        {"code":200,"data":["php","go"],"message":"welcome hg-gin page"}
    4. 查看access.log
        heige@daheige:/web/wwwlogs$ tailf  mygo.com-access.log
```
