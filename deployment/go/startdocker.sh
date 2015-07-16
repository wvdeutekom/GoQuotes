#! /bin/bash
. $(dirname $0)/config

echo "Starting up ${go_containername}"

docker run -d \
-p ${go_ports[0]} \
--name ${go_containername} \
${hub_username}/${go_hub_repository}:${go_imagename}
