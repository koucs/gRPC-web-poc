# Containerizing your Go Applications with Docker - Tutorial
# https://github.com/juaruipav/grpc-go-docker-helloworld

FROM golang:1.14

# Install grpc
RUN go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go

# Copy the project to be executed
RUN mkdir -p /go/src/greetclient 
COPY main.go /go/src/greetclient

ENTRYPOINT cd /go/src/greetclient && go run main.go
