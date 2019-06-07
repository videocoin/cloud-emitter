package service

import (
	"context"
	"fmt"
	"math/big"
	"time"

	sm "github.com/VideoCoin/cloud-pkg/streamManager"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/sirupsen/logrus"
)

type EventListenerConfig struct {
	StreamManager *sm.Manager
	Timeout       time.Duration
	Logger        *logrus.Entry
}

type EventListener struct {
	streamManager *sm.Manager
	timeout       time.Duration
	logger        *logrus.Entry
}

func NewEventListener(c *EventListenerConfig) *EventListener {
	return &EventListener{
		streamManager: c.StreamManager,
		timeout:       c.Timeout,
		logger:        c.Logger.WithField("component", "event-listener"),
	}
}

func (e *EventListener) LogStreamRequestEvent(ctx context.Context, streamID *big.Int, address common.Address) (chan *event, chan error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LogStreamRequestEvent")
	defer span.Finish()

	span.LogFields(
		log.String("address", address.Hex()),
		log.String("stream_id", streamID.String()),
	)

	addresses := []common.Address{address}
	streamIDs := []*big.Int{streamID}

	resultCh := make(chan *event, 1)
	errCh := make(chan error, 1)

	go func(context.Context) {
		span, _ := opentracing.StartSpanFromContext(ctx, "LogStreamRequestEventAsync")
		defer span.Finish()

		for timeout := time.After(e.timeout * time.Second); ; {
			select {
			case <-timeout:
				err := fmt.Errorf("failed to log stream request event and exit on timeout")
				e.logger.Error(err)
				span.LogFields(
					log.String("err", err.Error()),
				)

				errCh <- err
				break
			default:
				iterator, err := e.streamManager.FilterStreamRequested(
					new(bind.FilterOpts), addresses, streamIDs)
				if err != nil {
					err := fmt.Errorf("failed to log stream request event: %s", err.Error())
					e.logger.Error(err)
					span.LogFields(
						log.String("err", err.Error()),
					)

					errCh <- err
				}

				for {
					if iterator.Error() != nil {
						err := fmt.Errorf("failed to retrieve or parse log: %s", err.Error())
						e.logger.Error(err)
						span.LogFields(
							log.String("err", err.Error()),
						)

						errCh <- err
					}

					if iterator.Event != nil {
						e := iterator.Event

						span.LogFields(
							log.String("stream_id", e.StreamId.String()),
							log.String("address", e.Raw.Address.Hex()),
						)

						resultCh <- &event{
							Name:        EventStreamRequested,
							StreamID:    e.StreamId,
							Address:     e.Raw.Address,
							BlockNumber: e.Raw.BlockNumber,
							BlockHash:   e.Raw.BlockHash,
							TxHash:      e.Raw.TxHash,
							TxIndex:     e.Raw.TxIndex,
						}
						return
					}

					if !iterator.Next() {
						break
					}
				}

				time.Sleep(5 * time.Second)
			}
		}
	}(ctx)

	return resultCh, errCh
}

func (e *EventListener) LogStreamCreateEvent(ctx context.Context, streamID *big.Int) (chan *event, chan error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LogStreamCreateEvent")
	defer span.Finish()

	span.LogFields(
		log.String("stream_id", streamID.String()),
	)

	streamAddresses := []common.Address{}
	streamIDs := []*big.Int{streamID}

	resultCh := make(chan *event, 1)
	errCh := make(chan error, 1)

	go func(ctx context.Context) {
		span, _ := opentracing.StartSpanFromContext(ctx, "LogStreamCreateEventAsync")
		defer span.Finish()

		for timeout := time.After(e.timeout * time.Second); ; {
			select {
			case <-timeout:
				err := fmt.Errorf("failed to log stream create event and exit on timeout")
				e.logger.Error(err)
				span.LogFields(
					log.String("err", err.Error()),
				)

				errCh <- err
				break
			default:
				iterator, err := e.streamManager.FilterStreamCreated(
					new(bind.FilterOpts), streamAddresses, streamIDs)
				if err != nil {
					err := fmt.Errorf("failed to log stream create event: %s", err.Error())
					e.logger.Error(err)
					span.LogFields(
						log.String("err", err.Error()),
					)

					errCh <- err
				}

				for {
					if iterator.Error() != nil {
						err := fmt.Errorf("failed to retrieve or parse log: %s", err.Error())
						e.logger.Error(err)
						span.LogFields(
							log.String("err", err.Error()),
						)

						errCh <- err
					}

					if iterator.Event != nil {
						e := iterator.Event

						span.LogFields(
							log.String("stream_id", e.StreamId.String()),
							log.String("stream_address", e.StreamAddress.String()),
							log.String("address", e.Raw.Address.Hex()),
						)

						resultCh <- &event{
							Name:          EventStreamCreated,
							StreamID:      e.StreamId,
							Address:       e.Raw.Address,
							StreamAddress: e.StreamAddress,
							BlockNumber:   e.Raw.BlockNumber,
							BlockHash:     e.Raw.BlockHash,
							TxHash:        e.Raw.TxHash,
							TxIndex:       e.Raw.TxIndex,
						}
						return
					}

					if !iterator.Next() {
						break
					}
				}

				time.Sleep(5 * time.Second)
			}
		}
	}(ctx)

	return resultCh, errCh
}

func (e *EventListener) LogStreamApproveEvent(ctx context.Context, streamID *big.Int) (chan *event, chan error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LogStreamApproveEvent")
	defer span.Finish()

	span.LogFields(
		log.String("stream_id", streamID.String()),
	)

	streamIDs := []*big.Int{streamID}

	resultCh := make(chan *event, 1)
	errCh := make(chan error, 1)

	go func(context.Context) {
		span, _ := opentracing.StartSpanFromContext(ctx, "LogStreamApproveEventAsync")
		defer span.Finish()

		for timeout := time.After(e.timeout * time.Second); ; {
			select {
			case <-timeout:
				err := fmt.Errorf("failed to log stream approved event and exit on timeout")
				e.logger.Error(err)
				errCh <- err
				break
			default:
				iterator, err := e.streamManager.FilterStreamApproved(
					new(bind.FilterOpts), streamIDs)
				if err != nil {
					err := fmt.Errorf("failed to log stream approved event: %s", err.Error())
					e.logger.Error(err)
					span.LogFields(
						log.String("err", err.Error()),
					)

					errCh <- err
				}

				for {
					if iterator.Error() != nil {
						err := fmt.Errorf("failed to retrieve or parse log: %s", err.Error())
						e.logger.Error(err)
						span.LogFields(
							log.String("err", err.Error()),
						)

						errCh <- err
					}

					if iterator.Event != nil {
						e := iterator.Event

						span.LogFields(
							log.String("stream_id", e.StreamId.String()),
							log.String("address", e.Raw.Address.Hex()),
						)

						resultCh <- &event{
							Name:        EventStreamApproved,
							StreamID:    e.StreamId,
							Address:     e.Raw.Address,
							BlockNumber: e.Raw.BlockNumber,
							BlockHash:   e.Raw.BlockHash,
							TxHash:      e.Raw.TxHash,
							TxIndex:     e.Raw.TxIndex,
						}
						return
					}

					if !iterator.Next() {
						break
					}
				}

				time.Sleep(5 * time.Second)
			}
		}
	}(ctx)

	return resultCh, errCh
}

func (e *EventListener) LogInputChunkAddEvent(ctx context.Context, streamID *big.Int, chunkID *big.Int) (chan *event, chan error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LogInputChunkAddEvent")
	defer span.Finish()

	span.LogFields(
		log.String("stream_id", streamID.String()),
		log.String("chunk_id", chunkID.String()),
	)

	streamIDs := []*big.Int{streamID}
	chunkIDs := []*big.Int{chunkID}

	resultCh := make(chan *event, 1)
	errCh := make(chan error, 1)

	go func(context.Context) {
		span, _ := opentracing.StartSpanFromContext(ctx, "LogInputChunkAddEventAsync")
		defer span.Finish()

		for timeout := time.After(e.timeout * time.Second); ; {
			select {
			case <-timeout:
				err := fmt.Errorf("failed to log input chunk added event and exit on timeout")
				e.logger.Error(err)
				span.LogFields(
					log.String("err", err.Error()),
				)

				errCh <- err
				break
			default:
				iterator, err := e.streamManager.FilterInputChunkAdded(
					new(bind.FilterOpts), streamIDs, chunkIDs)
				if err != nil {
					err := fmt.Errorf("failed to log input chunk added event: %s", err.Error())
					e.logger.Error(err)
					span.LogFields(
						log.String("err", err.Error()),
					)

					errCh <- err
				}

				for {
					if iterator.Error() != nil {
						err := fmt.Errorf("failed to retrieve or parse log: %s", err.Error())
						e.logger.Error(err)
						span.LogFields(
							log.String("err", err.Error()),
						)

						errCh <- err
					}

					if iterator.Event != nil {
						e := iterator.Event

						span.LogFields(
							log.String("stream_id", e.StreamId.String()),
							log.String("chunk_id", e.ChunkId.String()),
							log.String("address", e.Raw.Address.Hex()),
						)

						resultCh <- &event{
							Name:        EventStreamInputChunkAdded,
							StreamID:    e.StreamId,
							ChunkID:     e.ChunkId,
							Address:     e.Raw.Address,
							BlockNumber: e.Raw.BlockNumber,
							BlockHash:   e.Raw.BlockHash,
							TxHash:      e.Raw.TxHash,
							TxIndex:     e.Raw.TxIndex,
						}
						return
					}

					if !iterator.Next() {
						break
					}
				}

				time.Sleep(5 * time.Second)
			}
		}
	}(ctx)

	return resultCh, errCh
}

func (e *EventListener) LogEndStreamEvent(ctx context.Context, streamID *big.Int, address common.Address) (chan *event, chan error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LogEndStreamEvent")
	defer span.Finish()

	span.LogFields(
		log.String("stream_id", streamID.String()),
		log.String("address", address.Hex()),
	)

	streamIDs := []*big.Int{streamID}
	caller := []common.Address{address}

	resultCh := make(chan *event, 1)
	errCh := make(chan error, 1)

	go func(context.Context) {
		span, _ := opentracing.StartSpanFromContext(ctx, "LogEndStreamEventAsync")
		defer span.Finish()

		for timeout := time.After(e.timeout * time.Second); ; {
			select {
			case <-timeout:
				err := fmt.Errorf("failed to log end stream event and exit on timeout")
				e.logger.Error(err)
				span.LogFields(
					log.String("err", err.Error()),
				)

				errCh <- err
				break
			default:
				iterator, err := e.streamManager.FilterStreamEnded(
					new(bind.FilterOpts), streamIDs, caller)
				if err != nil {
					err := fmt.Errorf("failed to log end stream event: %s", err.Error())
					e.logger.Error(err)
					span.LogFields(
						log.String("err", err.Error()),
					)

					errCh <- err
				}

				for {
					if iterator.Error() != nil {
						err := fmt.Errorf("failed to retrieve or parse log: %s", err.Error())
						e.logger.Error(err)
						span.LogFields(
							log.String("err", err.Error()),
						)

						errCh <- err
					}

					if iterator.Event != nil {
						e := iterator.Event

						span.LogFields(
							log.String("stream_id", e.StreamId.String()),
							log.String("address", e.Raw.Address.Hex()),
						)

						resultCh <- &event{
							Name:        EventStreamInputChunkAdded,
							StreamID:    e.StreamId,
							Address:     e.Raw.Address,
							BlockNumber: e.Raw.BlockNumber,
							BlockHash:   e.Raw.BlockHash,
							TxHash:      e.Raw.TxHash,
							TxIndex:     e.Raw.TxIndex,
						}
						return
					}

					if !iterator.Next() {
						break
					}
				}

				time.Sleep(5 * time.Second)
			}
		}
	}(ctx)

	return resultCh, errCh
}
