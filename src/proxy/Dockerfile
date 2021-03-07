FROM envoyproxy/envoy:v1.17.0

# for debug
# https://github.com/envoyproxy/envoy/blob/main/examples/front-proxy/Dockerfile-frontenvoy
RUN apt-get update && apt-get -q install -y curl

COPY envoy-grpc.yaml /etc/envoy/envoy.yaml

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /tmp/envoy_info.log
