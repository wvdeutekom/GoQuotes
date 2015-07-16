#! /bin/bash
. $(dirname $0)/../config

echo "Starting up ${go_containername}"

docker run -it \
-p ${go_ports[0]} \
--link ${rethink_containername}:rethinkdb \
--name ${go_containername} \
${hub_username}/${go_hub_repository}:${go_imagename}
