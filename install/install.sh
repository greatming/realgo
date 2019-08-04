#! /bin/bash

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


tar -zxf ./myframe.tar.gz 


sed -i " "  "s/myframe/${appName}/g"  `grep -rl  myframe ./myframe`

targetDir=$targetDir$appName
echo $targetDir

if [ -d $targetDir ]; then
    echo 'targetDir is exits'
    exit
fi

mv myframe $targetDir


if [ $? -ne 0 ];then
    echo "fail"
    exit
fi
echo "success"



