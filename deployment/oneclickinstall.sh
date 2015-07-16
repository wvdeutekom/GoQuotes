#! /bin/bash
. $(dirname $0)/config

basepath=${PWD}/$(dirname $0)/
echo "You may need to sign into Docker hub. If you have signed in already you can push enter"

sh ${basepath}/rethinkdb/startdocker.sh
sh ${basepath}/go/startdocker.sh
