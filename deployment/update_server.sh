#! /bin/bash
. ./config

# Pull new images from docker hub (that have been tested/verified by circleci)
pullimages () {
  docker pull ${hub_username}/${rethink_hub_repository}:${rethink_imagename}
  docker pull ${hub_username}/${go_hub_repository}:${go_imagename}
}

# Then stop and delete all containers
# Run installation script again to boot up all containers
pullimages && bash ./stopservices.sh 'yes'
bash ./oneclickinstall.sh
