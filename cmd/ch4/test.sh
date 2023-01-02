#!/usr/bin/env bash

#使用golang语言获取docker私有仓库的地址获取所有的镜像和tag信息，
#echo "test"
url="http://127.0.0.1:5000"
pwd=`pwd`
tes=`${pwd}/dockerImagesjson ${url}`
echo $tes
