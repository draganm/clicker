package proxy

import (
	"log"
	"net"
	"net/http"

	"github.com/draganm/bouncer"
	"github.com/draganm/clicker/comm"
	uuid "github.com/satori/go.uuid"
	"github.com/urfave/negroni"
)

func Proxy(bnd, remote, clickerServer string) error {
	c, err := net.Dial("udp", remote)
	if err != nil {
		return err
	}

	return bouncer.Proxy(bnd, remote, negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		id := uuid.NewV4()
		evt := comm.Event{
			Type:         "request",
			UUID:         id.String(),
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

		next.ServeHTTP(w, r)

	}))
}
