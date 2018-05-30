#!/bin/bash
#运维上线需执行该脚本
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

mkdir -p $root_dir/runtime/logs
chmod 777 -R $root_dir/runtime

#build ware
cd $root_dir/application/middleware
go install

#build controller
cd $root_dir/application/controller
go install

#build routes
cd $root_dir/application/routes
go install

#build app
cd $root_dir
go install
