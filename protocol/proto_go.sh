#!/usr/bin/env bash

#生成当前目录的proto文件
#protoc --proto_path=. --go_out=./protopb *.proto
#生成目录下的proto文件
currentDir=$(ls -l $(dirname $0) |awk '/^d/ {print $NF}')
for dir in ${currentDir}
do
    if [ ${dir} != "pb" ];then
        echo "generated " ${dir}
        #指定引入路径包含common
        #protoc --proto_path=${dir} --proto_path=./common --go_out=plugins=grpc:./pb ${dir}/*.proto
        protoc --proto_path=${dir} --go_out=plugins=grpc:./pb ${dir}/*.proto
    fi
done
