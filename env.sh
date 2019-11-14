#!/usr/bin/env bash
export GOPATH=$GOPATH:`pwd`

echo $GOPATH

export GOBIN=`pwd`/bin

echo $GOBIN

cd bin

rm -rf server

cd ..

rm -rf pkg

cd src

go build .

mv src ../bin/server

cd ../bin

./server