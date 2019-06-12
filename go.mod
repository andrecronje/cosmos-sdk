module github.com/andrecronje/cosmos-sdk

require (
	github.com/andrecronje/babble-abci latest
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.0.0-20190605094302-a0d1e3e36d50
	github.com/cosmos/cosmos-sdk latest
	github.com/cosmos/go-bip39 v0.0.0-20180618194314-52158e4697b8
	github.com/cosmos/ledger-cosmos-go v0.10.3
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.1
	github.com/gorilla/mux v1.7.0
	github.com/mattn/go-isatty v0.0.6=
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/rakyll/statik v0.1.4
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.0.3
	github.com/stretchr/testify v1.3.0
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/iavl v0.12.2
	github.com/tendermint/tendermint v0.31.7
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	google.golang.org/genproto v0.0.0-20180831171423-11092d34479b // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
