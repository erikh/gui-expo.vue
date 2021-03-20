// package pollsync is for updating lots of data from several sources and being
// able to poll it in one spot. The smallest of reactors
package pollsync

import (
	"context"
	"sync"
	"time"
)

type Instance struct {
	data         map[string]interface{}
	dataLock     sync.Mutex
	dataItemLock map[string]*sync.Mutex
	errChan      chan error
	context      context.Context
	cancelFunc   context.CancelFunc
}

func New(ctx context.Context) *Instance {
	ctx2, cancel := context.WithCancel(ctx)
	return &Instance{
		data:         map[string]interface{}{},
		dataItemLock: map[string]*sync.Mutex{},
		errChan:      make(chan error),
		context:      ctx2,
		cancelFunc:   cancel,
	}
}

func (i *Instance) inDataLock(source string, fun func()) {
	i.dataLock.Lock()
	var (
		srcLock *sync.Mutex
		ok      bool
	)

	srcLock, ok = i.dataItemLock[source]
	if !ok {
		srcLock = &sync.Mutex{}
		i.dataItemLock[source] = srcLock
	}
	i.dataLock.Unlock()
	srcLock.Lock()
	defer srcLock.Unlock()
	fun()
}

func (i *Instance) Data(source string) interface{} {
	var data interface{}
	i.inDataLock(source, func() {
		data = i.data[source]
	})
	return data
}

func (i *Instance) tickLoop(interval time.Duration, source string, fun func(context.Context) (interface{}, error)) {
	for range time.Tick(interval) {
		select {
		case <-i.context.Done():
			return
		default:
		}

		data, err := fun(i.context)
		if err != nil {
			i.errChan <- err
			return
		}

		i.inDataLock(source, func() {
			i.data[source] = data
		})
	}
}

func (i *Instance) Register(interval time.Duration, source string, fun func(context.Context) (interface{}, error)) error {
	var err error
	go i.tickLoop(interval, source, fun)
	return err
}

func (i *Instance) Error() error {
	select {
	case err := <-i.errChan:
		return err
	default:
		return nil
	}
}

func (i *Instance) Stop() {
	i.cancelFunc()
}
