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
#	tags=$(curl -s http://192.168.94.66:5000/v2/$im/tags/list | ./jq .tags)
#	echo $im
	tags=$(curl -s ${url}/v2/${im}/tags/list | ./jq-linux64 .tags)
#	if [[ -n ${tags} ]]; then
#		tags=${tags:1:-1}
#	else
#		tags=noTag
#	fi
	tags=${tags:1:-1}
	tag=${tags//\"}
	tag=${tag//\,}
#	echo $tag

	for t in $tag
	do
#		echo $i
		resultArr[$i]=$im:$t
		i=`expr $i + 1`;
	done
done

echo ${resultArr[*]}

