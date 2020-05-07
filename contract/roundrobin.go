package contract

import (
	"errors"
	"sync/atomic"
)

var ErrKSNotExists = errors.New("ks does not exist")

type RoundRobin interface {
	Next() *KSItem
}

type roundrobin struct {
	items []*KSItem
	next  uint32
}

func NewRoundRobin(items []*KSItem) (RoundRobin, error) {
	if len(items) == 0 {
		return nil, ErrKSNotExists
	}

	return &roundrobin{
		items: items,
	}, nil
}

func (r *roundrobin) Next() *KSItem {
	n := atomic.AddUint32(&r.next, 1)
	return r.items[(int(n)-1)%len(r.items)]
}
