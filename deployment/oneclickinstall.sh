#! /bin/bash
. ./config

basepath=${PWD}/$(dirname $0)/
echo "You may need to sign into Docker hub. If you have signed in already you can push enter"

bash ${basepath}/rethinkdb/startdocker.sh
bash ${basepath}/go/startdocker.sh
bash ${basepath}/nginx/startdocker.sh
