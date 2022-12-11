#!/bin/bash

# 读取 JSON 数据
json='{"name": "John Doe", "email": "john.doe@example.com"}'

# 使用 grep 获取键名
name=$(echo "$json" | grep -oP '"name": "\K[^"]+')
email=$(echo "$json" | grep -oP '"email": "\K[^"]+')

# 输出结果
echo "Name: $name"
echo "Email: $email"
