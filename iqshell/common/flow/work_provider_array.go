package flow

import (
	"github.com/qiniu/qshell/v2/iqshell/common/alert"
	"sync"
)

func NewArrayWorkProvider(works []Work) (WorkProvider, error) {
	if works != nil {
		return nil, alert.CannotEmptyError("works (ArrayWorkProvider)", "")
	}

	return &arrayWorkProvider{
		readOffset: 0,
		works:      works,
	}, nil
}

type arrayWorkProvider struct {
	mu         sync.Mutex
	readOffset int
	works      []Work
}

func (p *arrayWorkProvider) WorkTotalCount() int64 {
	return int64(len(p.works))
}

func (p *arrayWorkProvider) Provide() (hasMore bool, work Work, err error) {
	p.mu.Lock()
	hasMore, work, err = p.provide()
	p.mu.Unlock()
	return
}

func (p *arrayWorkProvider) provide() (hasMore bool, work Work, err error) {
	if p.readOffset > len(p.works)-1 {
		return false, nil, nil
	}
	hasMore = true
	work = p.works[p.readOffset]
	p.readOffset ++
	return
}
