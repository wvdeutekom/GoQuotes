#! /bin/bash
. $(dirname $0)/config

docker run -d \
-p ${rethink_ports[0]} \
-p ${rethink_ports[1]} \
-p ${rethink_ports[2]} \
--name ${rethink_containername} \
${hub_username}/${rethink_hub_repository}:${rethink_imagename}
