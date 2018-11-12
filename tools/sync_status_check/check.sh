#!/bin/bash


help() {
    echo "Usage: USER=<user> PASS=<passord> STAG=<stag_harbor_addr> PROD=<prod_harbor_addr> ./check.sh <repo> <tag>"
    exit 0
}

if [ $# -ne 2 ] ; then
    help
fi

repo=$1
tag=$2

./harborctl login -u $USER  -p $PASS --address $STAG
echo "-----"
#./harborctl job replication list -i 7 --status error -r $1 --address $STAG
#echo "-----"
./harborctl repository tag get -r $1 -t $2 --address $STAG

echo
echo
echo

echo "-----"
./harborctl login -u $USER -p $PASS --address $PROD
echo "-----"
./harborctl repository tag get -r $1 -t $2 --address $PROD
echo "-----"
./harborctl repository tag manifest -r $1 -t $2 -v v2 --address $PROD
echo "-----"
