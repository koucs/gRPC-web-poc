## About

[gRPC/Protocol Buffer Compiler Containers](https://github.com/namely/docker-protoc)

```shell
$ docker run --rm -v `pwd`:/defs namely/protoc-all -f echo.proto -l go
```

## Generate nodejs file

```shell
$ docker run --rm -v `pwd`:/defs namely/protoc-all -f helloworld.proto -l web
```

And you need to add `/* eslint-disable */` to .js files.

https://github.com/improbable-eng/grpc-web/issues/96#issuecomment-523448731

