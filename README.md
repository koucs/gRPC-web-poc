# gRPC-web-poc

POC of gRPC-web

### docker-compose command

```shell
$ docker-compose build --no-cache greetserver proxy
```

```
$ docker-compose up -d greetserver proxy
```

### docker command

```shell
$ docker network create -d bridge envoymesh
```

```shell
$ docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
8160a730d88d        bridge              bridge              local
b403ba23b223        envoymesh           bridge              local
1b6db2a29524        host                host                local
1827bd6b4431        none                null                local
```

```shell
$ docker run --rm -p 8080:8080 --network=envoymesh --network-alias proxy -it grpc-web-poc_proxy
```

```shell
$ docker run --rm -p 9090:9090 --network=envoymesh --network-alias greetserver -it grpc-web-poc_greetserver
```
#### React JS client App

```shell
$ npx create-react-app client --template typescript

$ cd client

$ npm install grpc-web google-protobuf
```

Refer [this](https://qiita.com/otanu/items/98d553d4b685a8419952).