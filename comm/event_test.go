package comm_test

import (
	"net/http"
	"testing"

	"github.com/draganm/clicker/comm"
	"github.com/stretchr/testify/assert"
)

func TestRequestEncodeDecode(t *testing.T) {
	req := comm.Event{
		Type:         "request",
		UUID:         "1234",
		Method:       "GET",
		RequestURI:   "/index.html",
		Header:       http.Header{},
		CapturedBody: []byte("TEST"),
	}

	data, err := req.Encode()
	assert.NoError(t, err)

	decoded, err := comm.Decode(data)
	assert.NoError(t, err)

	assert.Equal(t, req, decoded)
}
