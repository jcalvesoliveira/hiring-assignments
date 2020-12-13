FROM golang

COPY . /go/src/add
WORKDIR /go/src/add

# Downloading grpc_health_probe for K8s readinessProbe and livenessProbe
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

EXPOSE 50051

ENTRYPOINT go run cmd/service/service.go