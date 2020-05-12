#!/bin/bash

up(){
   ver="$1";
   file="$2";
   echo "------------->""$file";

   sed -E -i '' "s/github\.com\/vhaoran\/vchat .*/github\.com\/vhaoran\/vchat $ver/g"  "$file";
   echo "-------after update----------------"
   grep "vchat"  "$file"


};


VER=$(git ls-remote --heads |awk '{ print $1}')


up "$VER"  "$GOPATH/src/vchatintf/go.mod";
up "$VER"  "$GOPATH/src/vchatuser/go.mod";
up "$VER"  "$GOPATH/src/vchatutil/go.mod";
up "$VER"  "$GOPATH/src/vchatmsg/go.mod";
up "$VER"  "$GOPATH/src/vchatfinance/go.mod";
up "$VER"  "$GOPATH/src/vchatgw/go.mod";


cd "$GOPATH/src/vchatintf/" && go vet -composites=false ./...
cd "$GOPATH/src/vchatuser/" && go vet -composites=false ./...
cd "$GOPATH/src/vchatutil/" && go vet -composites=false ./...
cd  "$GOPATH/src/vchatmsg/" && go vet -composites=false ./...
cd  "$GOPATH/src/vchatfinance/" && go vet -composites=false ./...
cd "$GOPATH/src/vchatgw/" && go vet -composites=false ./...



echo "$VER"
