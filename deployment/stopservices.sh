#! /bin/bash
. ./config

echo "Stopping ${projectname} services.."


# Stop all the relevant docker containers
stopcontainers () {
  for i in $(docker ps -a | grep "${rethink_containername}" | cut -f1 -d" "); do
    echo "stopping $1"
    docker stop $i;
  done

  for i in $(docker ps -a | grep "${go_containername}" | cut -f1 -d" "); do
    echo "stopping $1"
    docker stop $i;
  done

}
confirm () {
  if [ "$1" = "yes" ]; then
    true
  else
    read -r -p "${1:-Delete containers too? [y/N]} " response
    case $response in
        [yY][eE][sS]|[yY])
            true
            ;;
        *)
            false
            ;;
    esac
  fi
}

# delete all the relevant docker containers (not the datavolume!)
deletecontainers () {
    for i in $(docker ps -a | grep ${rethink_containername} | grep -v ${rethink_datavolume} | cut -f1 -d" "); do
      echo "deleting $1"
      docker rm $i;
    done

    for i in $(docker ps -a | grep ${go_containername} | cut -f1 -d" "); do
      echo "deleting $1"
      docker rm $i;
    done
}

# Stop the containers and delete them, see above.
stopcontainers && confirm $1 && deletecontainers
