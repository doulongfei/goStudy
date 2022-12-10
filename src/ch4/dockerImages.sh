#!/usr/bin/env bash

url=http://127.0.0.1:5000
datas=$(curl -s ${url}/v2/_catalog | ./jq .repositories)
#echo $datas

datas=${datas:1:-1}

i=0
for image in $datas
do
#	echo $image
	im=${image//\"}
	im=${im//\,}
	tags=$(curl -s ${url}/v2/${im}/tags/list | ./jq .tags)
	tags=${tags:1:-1}
	tag=${tags//\"}
	tag=${tag//\,}
	for t in $tag
	do
		resultArr[$i]=$im:$t
		i=`expr $i + 1`;
	done
done

echo ${resultArr[*]}

