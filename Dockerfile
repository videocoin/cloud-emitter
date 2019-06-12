FROM golang:1.12.4 as builder
WORKDIR /go/src/github.com/videocoin/cloud-emitter
COPY . .
RUN make build

FROM bitnami/minideb:jessie
COPY --from=builder /go/src/github.com/videocoin/cloud-emitter/bin/emitter /opt/videocoin/bin/emitter
CMD ["/opt/videocoin/bin/emitter"]