#! /bin/bash
. ./config

echo "Stopping ${projectname} services.."


# Stop all the relevant docker containers
stopcontainers () {
    docker stop ${rethink_containername}
    docker stop ${go_containername}
    docker stop ${nginx_containername}
    for i in $(docker ps -a | grep "${rethink_containername}" | grep -v ${rethink_datavolume} | cut -f1 -d" "); do
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

    #Used to have a great for loop with grep, but this is more reliable
    #delete go container like this because the name is too generic
    docker rm ${go_containername} 
    docker rm ${rethink_containername}
    docker rm ${nginx_containername}
}

# Stop the containers and delete them, see above.
stopcontainers && confirm $1 && deletecontainers
