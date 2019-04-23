package service

import (
	pipelinesv1 "github.com/VideoCoin/cloud-api/pipelines/v1"
	"github.com/VideoCoin/cloud-pkg/mqmux"
	"github.com/sirupsen/logrus"
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

func (e *EventBus) UpdatePipelineStatus(req *pipelinesv1.UpdatePipelineRequest) error {
	e.logger.Infof("sending pipeline update: %v", req)
	return e.mq.Publish("pipeline/update", req)
}
