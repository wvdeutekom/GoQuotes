#! /bin/bash
. $(dirname $0)/config
echo ${deploy_dir}

basepath=${PWD}/$(dirname $0)/

buildimages () {
  echo "build images"
  docker build -t ${hub_username}/${rethink_hub_repository} ${basepath}/rethinkdb/.
  docker build -t ${hub_username}/${go_hub_repository} ${basepath}/../.
}

pushimages () {
  echo "push images"
  docker push ${hub_username}/${rethink_hub_repository}:${rethink_imagename}
  docker push ${hub_username}/${go_hub_repository}:${go_imagename}
}

rsyncfiles () {
  echo "rsync files"
  rsync -avzR ./deployment root@${deploy_server}:/${deploy_dir}
}

runupdatescript () {
  echo "run update script"
  ssh root@${deploy_server} "chmod -R +x ${deploy_dir}/deployment"
  ssh root@${deploy_server} "cd ${deploy_dir}/deployment; bash update_server.sh"
}

buildimages && pushimages && rsyncfiles && runupdatescript
