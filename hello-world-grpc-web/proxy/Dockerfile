FROM envoyproxy/envoy:latest

COPY ./envoy.yaml /etc/envoy/envoy.yaml

EXPOSE 50051

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml --component-log-level upstream:debug