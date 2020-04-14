#! /bin/bash

cd `dirname $0`

tmpdir=$GOPATH
OLD_IFS="$IFS"
IFS=":"
arr=($tmpdir)
IFS="$OLD_IFS"
targetDir=${arr[0]}/src/


appName=$1
if [ $# -lt 1 ];then
    appName=myframe
fi

if [ $# -gt 1 ];then
    targetDir=$2
fi


rm -rf  ./myframe

tar -zxf ./myframe.tar.gz 

sed -i ""  "s/myframe/${appName}/g"  `grep -rl  myframe ./myframe`

targetDir="${targetDir}/${appName}"


if [ -d $targetDir ]; then
    echo $targetDir 'folder is exits'
    exit
fi

mv myframe $targetDir


if [ $? -ne 0 ];then
    echo "fail"
    exit
fi
echo "success"



