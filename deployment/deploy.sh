#! /bin/bash
. $(dirname $0)/config

basepath=${PWD}/$(dirname $0)/
echo ${basepath}/../.

buildimages () {
  docker build -t ${hub_username}/${rethink_hub_repository} ${basepath}/rethinkdb/.
  docker build -t ${hub_username}/${go_hub_repository} ${basepath}/../.
}

pushimages () {
  docker push ${hub_username}/${rethink_hub_repository}:${rethink_imagename}
  docker push ${hub_username}/${go_hub_repository}:${go_imagename}
}

buildimages && pushimages


rsync -avzR ./deployment root@toaster.vdeute.com:/
      - ssh root@toaster.vdeute.com 'chmod -R +x /deployment'
      - ssh root@toaster.vdeute.com 'cd /deployment; sh update_server.sh'
