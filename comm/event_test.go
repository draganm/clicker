package comm_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/draganm/clicker/comm"
	"github.com/stretchr/testify/assert"
)

func TestRequestEncodeDecode(t *testing.T) {
	req := comm.Event{
		Time:                time.Now(),
		UUID:                "1234",
		Method:              "GET",
		RequestURI:          "/index.html",
		RequestHeader:       http.Header{},
		CapturedRequestBody: []byte("TEST"),
	}

	data, err := req.Encode()
	assert.NoError(t, err)

	decoded, err := comm.Decode(data)
	assert.NoError(t, err)

	assert.Equal(t, req, decoded)
}
