# pororift-backend
What the frontend/client calls for LoL data.



```
1. cd src; mkdir bin; export GOBIN=`pwd`/bin
2. go install ./...
3. touch api_key (Update)
4. ./bin/server
```


## API KEY
- Generate and save riot api key into `api_key` file in root project directory

## Static Content
- Static content need to be download using `updateData.sh`

```
git submodule update --init --recursive
```

# API
```
# Static art
http://localhost:3001/static/champion/tiles/Aatrox_0.jpg
http://localhost:3001/static/champion/splash/Aatrox_0.jpg
http://localhost:3001/static/champion/loading/Aatrox_0.jpg

# TODO: List of available art

```

















# Legacy API
```
localhost:3000/url

localhost:3000/summoner/:summonerIGN

localhost:3000/match/:matchid

http://localhost:3000/champ_by_key/1

http://localhost:3000/championList

http://localhost:3000/champ_icon/1

# Champion Rotation
http://localhost:3000/champ_rotation

```

# Static Images
```
# Get what images availables
# top paths are: loading, splash, tiles

http://localhost:3000/images

# Example: get Image
http://localhost:3000/champion/splash/Aatrox_1.jpg

```


# Docker Development (For Dev use only)

## Install Docker
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-16-04

BUILD
```
## docker build -t docker_node .
docker-compose up --build -d
```

Create container with mounted current directory
```
docker container run -dit -p 3000:3000 -v $(pwd):/usr/workspace --name dev_env ubuntu_image
```

RUN in interactive
```
docker container exec -it dev bash
nodemon --exec ./node_modules/.bin/babel src/index.js --ignore data/
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

docker rm $(docker ps -aq) # -f remove all docker containers
docker rmi $(docker images -q) # Remove all images


```
