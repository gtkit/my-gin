#!/bin/bash
RootPath=$2
echo "ssh://git@gitlab.superjq.com:47383${PROJECT_PATH}"

echo ${RootPath}

if [ -n "$1" ]
then
    Branch=$1
else
    Branch=master
fi


## 第一次部署代码
if [ ! -d "${RootPath}" ];then
	 echo "ssh://git@gitlab.superjq.com:47383${PROJECT_PATH}"
   echo "RootPath:"$RootPath
   git clone -b ${Branch} ssh://git@gitlab.superjq.com:47383${PROJECT_PATH}.git $RootPath
fi

cd $RootPath
echo "RootPath:"$RootPath
## 拉取代码
git clean -f -d > /dev/null
git fetch origin
git reset --hard origin/${Branch}

go env -w GOPROXY=https://goproxy.cn,direct


supervisorctl restart go-assistant-desktop > /dev/null 2>&1
echo "部署完成！"
