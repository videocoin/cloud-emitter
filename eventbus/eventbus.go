package eventbus

import (
	"context"
	"encoding/json"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	accountsv1 "github.com/videocoin/cloud-api/accounts/v1"
	v1 "github.com/videocoin/cloud-api/emitter/v1"
	faucetcli "github.com/videocoin/cloud-pkg/faucet"
	"github.com/videocoin/cloud-pkg/mqmux"
	tracerext "github.com/videocoin/cloud-pkg/tracer"
)

type Config struct {
	Logger *logrus.Entry
	URI    string
	Name   string
	Faucet *faucetcli.Client
}

type EventBus struct {
	logger *logrus.Entry
	mq     *mqmux.WorkerMux
	faucet *faucetcli.Client
}

func New(c *Config) (*EventBus, error) {
	mq, err := mqmux.NewWorkerMux(c.URI, c.Name)
	if err != nil {
		return nil, err
	}
	return &EventBus{
		logger: c.Logger,
		faucet: c.Faucet,
		mq:     mq,
	}, nil
}

func (e *EventBus) Start() error {
	err := e.mq.Consumer("accounts.events", 1, false, e.handleAccountEvent)
	if err != nil {
		return err
	}

	err = e.mq.Publisher("emitter.events")
	if err != nil {
		return err
	}

	return e.mq.Run()
}

func (e *EventBus) Stop() error {
	return e.mq.Close()
}

func (e *EventBus) handleAccountEvent(d amqp.Delivery) error {
	var span opentracing.Span
	tracer := opentracing.GlobalTracer()
	spanCtx, err := tracer.Extract(opentracing.TextMap, mqmux.RMQHeaderCarrier(d.Headers))

	e.logger.Debugf("handling body: %+v", string(d.Body))

	if err != nil {
		span = tracer.StartSpan("eventbus.handleAccountEvent")
	} else {
		span = tracer.StartSpan("eventbus.handleAccountEvent", ext.RPCServerOption(spanCtx))
	}

	defer span.Finish()

	req := new(accountsv1.Event)
	err = json.Unmarshal(d.Body, req)
	if err != nil {
		tracerext.SpanLogError(span, err)
		return err
	}

	span.SetTag("event_type", req.Type.String())
	span.SetTag("user_id", req.UserID)
	span.SetTag("address", req.Address)

	logger := e.logger.WithFields(logrus.Fields{
		"event_type": req.Type.String(),
		"user_id":    req.UserID,
		"address":    req.Address,
	})
	logger.Debugf("handling request %+v", req)

	switch req.Type {
	case accountsv1.EventTypeAccountCreated:
		{
			err := e.EmitAccountCreated(context.Background(), req.UserID, req.Address)
			if err != nil {
				e.logger.WithField("address", req.Address).Errorf("failed to emit account created")
			}

			err = e.faucet.Do(req.Address, 100)
			if err != nil {
				e.logger.WithField("address", req.Address).Errorf("failed to faucet")
				return nil
			}
		}
	case accountsv1.EventTypeUnknown:
		e.logger.Error("event type is unknown")
	}

	return nil
}

func (e *EventBus) EmitAccountCreated(ctx context.Context, userID, address string) error {
	headers := make(amqp.Table)

	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		ext.SpanKindRPCServer.Set(span)
		ext.Component.Set(span, "emitter")
		err := span.Tracer().Inject(
			span.Context(),
			opentracing.TextMap,
			mqmux.RMQHeaderCarrier(headers),
		)
		if err != nil {
			e.logger.Errorf("failed to span inject: %s", err)
		}
	}
	event := &v1.Event{
		Type:    v1.EventTypeAccountCreated,
		UserID:  userID,
		Address: address,
	}

	err := e.mq.PublishX("emitter.events", event, headers)
	if err != nil {
		return err
	}

	return nil
}
