module github.com/videocoin/cloud-emitter

go 1.14

require (
	github.com/AlekSi/pointer v1.1.0
	github.com/ethereum/go-ethereum v1.9.13
	github.com/gogo/protobuf v1.3.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.5.0
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/videocoin/cloud-api v0.2.15
	github.com/videocoin/cloud-pkg v0.0.6
	github.com/videocoin/go-protocol v0.0.8-0.20200519072212-ad37943377e7
	github.com/videocoin/go-staking v0.3.0
	google.golang.org/grpc v1.29.1
)

replace github.com/videocoin/cloud-api => ../cloud-api

replace github.com/videocoin/cloud-pkg => ../cloud-pkg
