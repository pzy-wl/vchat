#!/bin/bash

up(){
   ver="$1";
   file="$2";
   echo " path is --->""$file";

   sed -E -i '' "s/github\.com\/vhaoran\/vchat .*/github\.com\/vhaoran\/vchat $ver/g"  "$file";
};


VER=$(git ls-remote --heads |awk '{ print $1}')


up "$VER"  "$GOPATH/src/vchatuser/go.mod";
up "$VER"  "$GOPATH/src/vchatutil/go.mod";
up "$VER"  "$GOPATH/src/vchatmsg/go.mod";
up "$VER"  "$GOPATH/src/vchatfinance/go.mod";
up "$VER"  "$GOPATH/src/vchatgw/go.mod";



echo "$VER"
