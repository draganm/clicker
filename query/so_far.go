package query

import (
	"log"
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
	"github.com/draganm/clicker/comm"
	"github.com/draganm/zathras/topic"
)

type WebStatsBucket struct {
	Start         time.Time
	Duration      time.Duration
	Requests      int
	BytesReceived int
	BytesSent     int
	Errors        int
}

type SoFar struct {
	sync.Mutex
	stats                 *WebStatsBucket
	closeSubscriptionChan chan interface{}
	*emission.Emitter
}

func NewSoFar(t *topic.Topic) *SoFar {
	// buckets := []*WebStatsBucket{}
	// for i := 0; i < 5*60; i++ {
	// 	buckets = append(buckets, &WebStatsBucket{})
	// }
	evtChan, closeChan := t.Subscribe(0)
	// <-chan Event, chan interface{}

	sf := &SoFar{
		stats: &WebStatsBucket{},
		closeSubscriptionChan: closeChan,
		Emitter:               emission.NewEmitter(),
	}
	go func() {
		for evt := range evtChan {
			webEvt, err := comm.Decode(evt.Data)
			if err != nil {
				log.Println(err)
			} else {
				sf.webEvent(webEvt)
			}

		}
	}()

	return sf

}

func (s *SoFar) AddUpdateListener(fn func(WebStatsBucket)) {
	s.Lock()
	defer s.Unlock()
	s.AddListener("update", fn)
	go fn(*s.stats)
}

func (s *SoFar) RemoveUpdateListener(fn func(WebStatsBucket)) {
	s.Lock()
	defer s.Unlock()
	s.RemoveListener("update", fn)
}

func (s *SoFar) webEvent(evt comm.Event) {
	s.Lock()
	defer s.Unlock()
	s.stats.BytesReceived += evt.BytesRead
	s.stats.BytesSent += evt.BytesWritten
	s.stats.Requests++
	if evt.StatusCode >= 400 {
		s.stats.Errors++
	}
	s.Emit("update", *s.stats)
}

func (s *SoFar) Close() {
	close(s.closeSubscriptionChan)
}
