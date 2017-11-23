export GOROOT=/usr/local/go
export GOPATH=/root/gopath
export GOBIN=/root/gopath/bin
export LD_LIBRARY_PATH=/root/bip44cxx

rm -fr hdwallet
go build -o hdwallet ./
./hdwallet
# go run gowallet.go -n 3 -v BTC
# go run gowallet.go -n 3 
# -lbitcoin -lbitcoin-client -lbitcoin_server -lbitcoin_consensus -lbitcoin_crypto -lstdc++ -lboost_filesystem -lboost_system -lboost_thread -lm -lsecp256k1
# include "interface2.h"