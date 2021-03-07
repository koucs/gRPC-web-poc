# gRPC-web-poc
POC of gRPC-web

### docker-compose command

```shell
$ docker-compose build --no-cache greetserver
```

```
$ docker-compose up -d greetserver
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
$ docker run --rm -p 8080:8080 --network=envoymesh -it grpc-web-poc_proxy
```

```shell
$ docker run --rm -p 9090:9090 --network=envoymesh --network-alias server -it grpc-web-poc_server
```

```shell
$ docker run --rm -p 9091:9091 --network=envoymesh --network-alias greetserver -it grpc-web-poc_proxy
```
