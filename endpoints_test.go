package pubnub

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeEndpointOpts struct {
	pubnub *PubNub
}

func (o *fakeEndpointOpts) buildPath() (string, error) {
	return "/my/path", nil
}

func (o *fakeEndpointOpts) buildQuery() (*url.Values, error) {
	q := &url.Values{}

	q.Set("a", "2")

	q.Set("b", "hey")

	return q, nil
}

func (o *fakeEndpointOpts) buildBody() ([]byte, error) {
	return []byte("myBody"), nil
}

func (o *fakeEndpointOpts) jobQueue() chan *JobQItem {
	return o.pubnub.jobQueue
}

func (o *fakeEndpointOpts) config() Config {
	return *o.pubnub.Config
}

func (o *fakeEndpointOpts) client() *http.Client {
	return o.pubnub.GetClient()
}

func (o *fakeEndpointOpts) validate() error {
	return nil
}

func (o *fakeEndpointOpts) context() Context {
	return o.context()
}

func (o *fakeEndpointOpts) httpMethod() string {
	return "GET"
}

func (o *fakeEndpointOpts) operationType() OperationType {
	return PNSubscribeOperation
}

func (o *fakeEndpointOpts) telemetryManager() *TelemetryManager {
	return o.pubnub.telemetryManager
}

func xTestBuildURL(t *testing.T) {
	assert := assert.New(t)

	pnc := NewConfig()
	pn := NewPubNub(pnc)

	e := &fakeEndpointOpts{
		pubnub: pn,
	}

	url, err := buildURL(e)

	assert.Nil(err)
	assert.Equal("https://ps.pndsn.com/my/path?a=2&b=hey", url.RequestURI())
}
