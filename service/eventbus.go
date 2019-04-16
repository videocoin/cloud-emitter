package service

import (
	streamsv1 "github.com/VideoCoin/cloud-api/streams/v1"
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
	err := e.mq.Publisher("stream/update-status")
	if err != nil {
		return err
	}

	err = e.mq.Publisher("stream/update-stream-address")
	if err != nil {
		return err
	}

	return nil
}

func (e *EventBus) registerConsumers() error {
	return nil
}

func (e *EventBus) UpdateStreamStatus(req *streamsv1.UpdateStreamRequest) error {
	return e.mq.Publish("stream/update-status", req)
}

func (e *EventBus) UpdateStreamAddress(req *streamsv1.UpdateStreamRequest) error {
	return e.mq.Publish("stream/update-stream-address", req)
}
