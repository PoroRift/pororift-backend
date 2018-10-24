# pororift-backend
What the frontend/client calls for LoL data.














# Docker Development 
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