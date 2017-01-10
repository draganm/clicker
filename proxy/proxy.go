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

		ww := newWriterWrapper(w, 1024)
		readerWrapper := newBodyReaderWrapper(r.Body, 1024)
		r.Body = readerWrapper
		next.ServeHTTP(ww, r)
		{
			evt := comm.Event{
				Time:                 time.Now(),
				UUID:                 id,
				RequestHeader:        r.Header,
				ResponseHeader:       w.Header(),
				Method:               r.Method,
				RequestURI:           r.RequestURI,
				StatusCode:           ww.statusCode,
				CapturedResponseBody: ww.capturedData(),
				CapturedRequestBody:  readerWrapper.capturedData(),
				BytesRead:            readerWrapper.bytesRead,
				BytesWritten:         ww.bytesWritten,
				LastByteReadAt:       readerWrapper.lastReadAt,
				LastByteWrittenAt:    ww.lastWrittenAt,
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
