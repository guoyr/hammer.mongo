#!/bin/bash

# how many DB to be insert into
export HT_MULTI_DB=1

# how many collections per DB
export HT_MULTI_COLLECTION=16

# number of workers
export HT_CMD_WORKERS=16

# report stats every N second
export HT_CMD_MONITOR_INTERVAL=10

# how long shall the run in sercond, 0 means infinity
export HT_CMD_TOTAL_TIME=40

# how long shall the run in number of OP, 0 means infinity
export HT_CMD_TOTAL_OPS=0

# bulk insert size per batch
export HT_INSERT_BATCH_SIZE=1000

# server URL
export HT_SERVER_URL="localhost:27017"

# check hammer binary
PLATFORM='unknown'
unamestr=`uname`
if [[ "$unamestr" == 'Linux' ]]; then
    PLATFORM='linux'
elif [[ "$unamestr" == 'Darwin' ]]; then
    PLATFORM='macos'
fi

BINARY="godd run ../hammer.mongo.go"
if [ -x "../bin/hammer.$PLATFORM" ]; then
    echo "Found executable binary"
    BINARY="../bin/hammer.$PLATFORM" 
fi

$BINARY-profile=BULKINSERT -max -worker $HT_CMD_WORKERS -server $HT_SERVER_URL -monitor $HT_CMD_MONITOR_INTERVAL -total $HT_CMD_TOTAL_OPS -totaltime $HT_CMD_TOTAL_TIME

