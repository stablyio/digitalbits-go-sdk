package server

import (
	"math/big"
	"net/http"

	"github.com/digitalbitsorg/go/services/bifrost/bitcoin"
	"github.com/digitalbitsorg/go/services/bifrost/config"
	"github.com/digitalbitsorg/go/services/bifrost/database"
	"github.com/digitalbitsorg/go/services/bifrost/ethereum"
	"github.com/digitalbitsorg/go/services/bifrost/queue"
	"github.com/digitalbitsorg/go/services/bifrost/sse"
	"github.com/digitalbitsorg/go/services/bifrost/stellar"
	"github.com/digitalbitsorg/go/support/log"
)

// ProtocolVersion is the version of the protocol that Bifrost server and
// JS SDK use to communicate.
const ProtocolVersion int = 1

type Server struct {
	BitcoinListener            *bitcoin.Listener            `inject:""`
	BitcoinAddressGenerator    *bitcoin.AddressGenerator    `inject:""`
	Config                     *config.Config               `inject:""`
	Database                   database.Database            `inject:""`
	EthereumListener           *ethereum.Listener           `inject:""`
	EthereumAddressGenerator   *ethereum.AddressGenerator   `inject:""`
	StellarAccountConfigurator *stellar.AccountConfigurator `inject:""`
	TransactionsQueue          queue.Queue                  `inject:""`
	SSEServer                  sse.ServerInterface          `inject:""`

	MinimumValueBtc string
	MinimumValueEth string

	minimumValueSat int64
	minimumValueWei *big.Int
	httpServer      *http.Server
	log             *log.Entry
}

type GenerateAddressResponse struct {
	ProtocolVersion int    `json:"protocol_version"`
	Chain           string `json:"chain"`
	Address         string `json:"address"`
}
