package pubnub

import (
	"fmt"
	"net/url"
	"testing"

	h "github.com/pubnub/go/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func init() {
	pnconfig = NewConfig()

	pnconfig.PublishKey = "pub_key"
	pnconfig.SubscribeKey = "sub_key"
	pnconfig.SecretKey = "secret_key"

	pubnub = NewPubNub(pnconfig)
}

func TestGrantRequestBasic(t *testing.T) {
	assert := assert.New(t)

	opts := &GrantOpts{
		AuthKeys: []string{"my-auth-key"},
		Channels: []string{"ch"},
		Groups:   []string{"cg"},
		Read:     true,
		Write:    true,
		Manage:   true,
		Ttl:      "5000",
		pubnub:   pubnub,
	}

	path, err := opts.buildPath()
	assert.Nil(err)
	u := &url.URL{
		Path: path,
	}

	h.AssertPathsEqual(t,
		fmt.Sprintf("/v1/auth/grant/sub-key/%s", opts.pubnub.Config.SubscribeKey),
		u.EscapedPath(), []int{})

	query, err := opts.buildQuery()
	assert.Nil(err)

	expected := &url.Values{}
	expected.Set("auth", "my-auth-key")
	expected.Set("channel", "ch")
	expected.Set("channel-group", "cg")
	expected.Set("r", "1")
	expected.Set("w", "1")
	expected.Set("m", "1")
	expected.Set("ttl", "5000")
	h.AssertQueriesEqual(t, expected, query,
		[]string{"pnsdk", "uuid", "timestamp"}, []string{})

	body, err := opts.buildBody()
	assert.Nil(err)
	assert.Equal([]byte{}, body)
}
