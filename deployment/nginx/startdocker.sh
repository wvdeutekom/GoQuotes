#! /bin/bash
. $(dirname $0)/../config

echo "here wo go folks: ${nginx_hub_username}/${ngxinx_hub_reopository}:${nginx_imagename}"
docker run -d \
-p ${nginx_port[0]} \
-v /var/run/docker.sock:/tmp/docker.sock:ro \
--name ${nginx_containername} \
${nginx_hub_username}/${ngxinx_hub_reopository}:${nginx_imagename}
