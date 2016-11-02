#!/bin/bash

APP_NAME="iax.bidtimes"

cd `dirname $0`
BIN_DIR=`pwd`

pid=`ps aux|grep -v grep | grep -v "$BIN_DIR" | grep "$APP_NAME" | awk '{print $2}'`

if [ -n "$pid" ] && [ "$pid" > 0 ];
then 
	echo "kill $pid"
	kill "$pid"
fi

cd ..
nohup ./"$APP_NAME" > ./"$APP_NAME".stdout.log 2>&1 &

sleep 0.06

ps aux|grep -v grep | grep -v "$BIN_DIR" | grep "$APP_NAME"