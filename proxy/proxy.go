package proxy

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/draganm/bouncer"
	"github.com/draganm/clicker/comm"
	uuid "github.com/satori/go.uuid"
	"github.com/urfave/negroni"
)

func Proxy(bnd, remote, clickerServer string) error {
	c, err := net.Dial("udp", clickerServer)
	if err != nil {
		return err
	}

	return bouncer.Proxy(bnd, remote, negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		id := uuid.NewV4().String()
		{
			evt := comm.Event{
				Type:         "request",
				Time:         time.Now(),
				UUID:         id,
				Header:       r.Header,
				Method:       r.Method,
				RequestURI:   r.RequestURI,
				CapturedBody: []byte{},
			}

			data, err := evt.Encode()
			if err != nil {
				log.Println(err)

			} else {
				c.Write(data)
			}
		}
		next.ServeHTTP(w, r)
		{
			evt := comm.Event{
				Type:   "response",
				Time:   time.Now(),
				UUID:   id,
				Header: w.Header(),
				// Method:       r.Method,
				// RequestURI:   r.RequestURI,
				CapturedBody: []byte{},
			}

			data, err := evt.Encode()
			if err != nil {
				log.Println(err)

			} else {
				c.Write(data)
			}
		}

	}))
}
