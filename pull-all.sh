#!/bin/bash

fn(){
   
   DIR="$1"
   echo "-------->-""$DIR""----"
   cd $DIR
   git pull origin master

};




fn   "$GOPATH/src/vchatintf/";
fn   "$GOPATH/src/vchatuser/";
fn   "$GOPATH/src/vchatutil/";
fn   "$GOPATH/src/vchatmsg/";
fn   "$GOPATH/src/vchatfinance/";
fn  "$GOPATH/src/vchatgw/";

fn   "$GOPATH/src/intfuser/";
fn   "$GOPATH/src/intfgw/";




echo "------complete------------------"
