#!/bin/sh

#mydir=serialization-bench/gogopb
curdir=`pwd`
GOSRC=$GOPATH/src/
res=$(python -c "import os.path; print os.path.relpath('$curdir', '$GOSRC')")

cd $GOSRC
protoc --gogo_out=. $res/*.proto
