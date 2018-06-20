package stellartoml

import (
	"net/http"
	"strings"
	"testing"

	"github.com/digitalbitsorg/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientURL(t *testing.T) {
	//HACK:  we're testing an internal method rather than setting up a http client
	//mock.

	c := &Client{UseHTTP: false}
	assert.Equal(t, "https://digitalbits.io/.well-known/digitalbits.toml", c.url("digitalbits.io"))

	c = &Client{UseHTTP: true}
	assert.Equal(t, "http://digitalbits.io/.well-known/digitalbits.toml", c.url("digitalbits.io"))
}

func TestClient(t *testing.T) {
	h := httptest.NewClient()
	c := &Client{HTTP: h}

	// happy path
	h.
		On("GET", "https://digitalbits.io/.well-known/digitalbits.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation"`,
		)
	stoml, err := c.GetStellarToml("digitalbits.io")
	require.NoError(t, err)
	assert.Equal(t, "https://localhost/federation", stoml.FederationServer)

	// stellar.toml exceeds limit
	h.
		On("GET", "https://toobig.org/.well-known/digitalbits.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation`+strings.Repeat("0", StellarTomlMaxSize)+`"`,
		)
	stoml, err = c.GetStellarToml("toobig.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "digitalbits.toml response exceeds")
	}

	// not found
	h.
		On("GET", "https://missing2.org/.well-known/stellar.toml").
		ReturnNotFound()
	stoml, err = c.GetStellarToml("missing2.org")
	assert.EqualError(t, err, "http request failed with non-200 status code")

	// invalid toml
	h.
		On("GET", "https://json2.org/.well-known/stellar.toml").
		ReturnJSON(http.StatusOK, map[string]string{"hello": "world"})
	stoml, err = c.GetStellarToml("json2.org")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "toml decode failed")
	}
}
