#! /bin/bash
str=""
if [ $str ]; then
  echo $str,不为null
else
  echo $str,为null
fi

s1="null"
if [ $s1 == "null" ]; then
  echo $1,为null
else
  echo $s1,bu为null
fi