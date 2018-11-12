#!/bin/bash


help() {
    echo "Usage: USER=<user> PASS=<passord> ./check.sh <repo> <tag>"
    exit 0
}

if [ $# -ne 2 ] ; then
    help
fi

user=$USER
pass=$PASS

repo=$1
tag=$2

./harborctl login -u $user  -p $pass --address stag-reg.llsops.com
echo "-----"
#./harborctl job replication list -i 7 --status error -r $1 --address stag-reg.llsops.com
#echo "-----"
./harborctl repository tag get -r $1 -t $2 --address stag-reg.llsops.com

echo
echo
echo

echo "-----"
./harborctl login -u $user -p $pass --address prod-reg.llsops.com
echo "-----"
./harborctl repository tag manifest -r $1 -t $2 -v v2 --address prod-reg.llsops.com
echo "-----"
./harborctl repository tag get -r $1 -t $2 --address prod-reg.llsops.com
echo "-----"
