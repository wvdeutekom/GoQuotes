#! /bin/bash
. $(dirname $0)/../config

echo "Starting up ${rethink_containername}"

find=$(docker ps -a | grep ${rethink_datavolume} | awk '{ print $7;}')
if [ -z $find ]; then
  echo 'creating missing datavolume';
  docker create -v /data --name ${rethink_datavolume} ${hub_username}/${rethink_hub_repository}:${rethink_imagename}
fi

docker run -d \
-p ${rethink_ports[0]} \
-p ${rethink_ports[1]} \
-p ${rethink_ports[2]} \
--volumes-from ${rethink_datavolume} \
--name ${rethink_containername} \
${hub_username}/${rethink_hub_repository}:${rethink_imagename}
