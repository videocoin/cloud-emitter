FROM alpine:3.7

COPY bin/vc_emitter /opt/videocoin/bin/vc_emitter

CMD ["/opt/videocoin/bin/vc_emitter"]
