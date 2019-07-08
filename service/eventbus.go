package service

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	pipelinesv1 "github.com/videocoin/cloud-api/pipelines/v1"
	"github.com/videocoin/cloud-pkg/mqmux"
)

type EventBus struct {
	logger *logrus.Entry
	mq     *mqmux.WorkerMux
}

func NewEventBus(mq *mqmux.WorkerMux, logger *logrus.Entry) (*EventBus, error) {
	return &EventBus{
		logger: logger,
		mq:     mq,
	}, nil
}

func (e *EventBus) Start() error {
	err := e.registerPublishers()
	if err != nil {
		return err
	}

	err = e.registerConsumers()
	if err != nil {
		return err
	}

	return e.mq.Run()
}

func (e *EventBus) Stop() error {
	return e.mq.Close()
}

func (e *EventBus) registerPublishers() error {
	err := e.mq.Publisher("pipeline/update")
	if err != nil {
		return err
	}

	return nil
}

func (e *EventBus) registerConsumers() error {
	return nil
}

func (e *EventBus) UpdatePipelineStreamStatus(span opentracing.Span, req *pipelinesv1.UpdatePipelineStreamRequest) error {
	e.logger.Infof("sending pipeline stream update: %v", req)

	headers := make(amqp.Table)

	ext.SpanKindRPCServer.Set(span)
	ext.Component.Set(span, "pipelines")

	span.Tracer().Inject(
		span.Context(),
		opentracing.TextMap,
		mqmux.RMQHeaderCarrier(headers),
	)

	return e.mq.PublishX("pipeline/update", req, headers)
}
