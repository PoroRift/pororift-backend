# pororift-backend
What the frontend/client calls for LoL data.


# Docker Development 

## Install Docker
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-16-04

BUILD
```
docker build -t ubuntu_image .
```

Create container with mounted current directory
```
docker container run -dit -p 3000:3000 -v $(pwd):/usr/workspace --name dev_env ubuntu_image
```

RUN in interactive
```
docker container exec -it dev_env bash
```

Access
```
localhost:3000
```

## Clean up
show all docker instances
```
docker container ps -a
```

show all containers
```
container ls -a
```

Removal
```
docker image rm <ID/name>  # remove docker image
docker container rm <ID/name> # remove docker container

docker rm $(docker ps -aq) -f remove all docker containers
```