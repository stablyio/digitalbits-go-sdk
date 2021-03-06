package address

import (
	"testing"

	"github.com/digitalbitsorg/go/support/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cases := []struct {
		Name            string
		Domain          string
		ExpectedAddress string
	}{
		{"tom", "digitalbits.io", "tom*digitalbits.io"},
		{"", "digitalbits.io", "*digitalbits.io"},
		{"tom", "", "tom*"},
	}

	for _, c := range cases {
		actual := New(c.Name, c.Domain)
		assert.Equal(t, c.ExpectedAddress, actual)
	}
}

func TestSplit(t *testing.T) {
	cases := []struct {
		CaseName       string
		Address        string
		ExpectedName   string
		ExpectedDomain string
		ExpectedError  error
	}{
		{"happy path", "tom*digitalbits.io", "tom", "digitalbits.io", nil},
		{"blank", "", "", "", ErrInvalidAddress},
		{"blank name", "*digitalbits.io", "", "", ErrInvalidName},
		{"blank domain", "tom*", "", "", ErrInvalidDomain},
		{"invalid domain", "tom*--3.com", "", "", ErrInvalidDomain},
	}

	for _, c := range cases {
		name, domain, err := Split(c.Address)

		if c.ExpectedError == nil {
			assert.Equal(t, c.ExpectedName, name)
			assert.Equal(t, c.ExpectedDomain, domain)
		} else {
			assert.Equal(t, c.ExpectedError, errors.Cause(err))
		}
	}
}
