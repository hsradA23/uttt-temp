# Ultimate Tic Tac Toe game

### Building the binary
```
./build
```

### Running the app
There are 2 docker compose configs, one can be used to run locally with CLoudflared Tunnel,
the other can be used to host normally, exposing port 5000
```
docker-compose -f deploy/docker-compose.yaml up
```
Cloudflared logs will give the URL that you need to access to play the game with other people.

### Building the docker image without running
```
docker-compose -f deploy/docker-compose.yaml build
```
