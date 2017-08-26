// Package containing tendermints basecoin plugin
package plugin

import (
	"fmt"

	abci "github.com/tendermint/abci/types"
	"github.com/tendermint/basecoin/types"
	bptypes "github.com/torresashjian/blockproc/types"
	//"github.com/tendermint/go-wire"
)

type State struct {
	Counter   int
	TotalFees types.Coins
}

type CounterTx struct {
	Valid bool
	Fee   types.Coins
}

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Name() string {
	return "blockproc"
}

func (cp *App) StateKey() []byte {
	return []byte(fmt.Sprintf("BlockProc.State"))
}

func (cp *App) SetOption(store types.KVStore, key, value string) (log string) {
	return ""
}

func (cp *App) RunTx(store types.KVStore, ctx types.CallContext, txBytes []byte) (res abci.Result) {
	fmt.Println("In Plugin RunTx")

	// Decode tx
	var tx bptypes.Step
	err := wire.ReadBinaryBytes(txBytes, &tx)
	if err != nil {
		return abci.ErrBaseEncodingError.AppendLog("Error decoding tx: " + err.Error()).PrependLog("CounterTx Error: ")
	}

	return abci.OK
}

func (cp *App) InitChain(store types.KVStore, vals []*abci.Validator) {
}

func (cp *App) BeginBlock(store types.KVStore, hash []byte, header *abci.Header) {
}

func (cp *App) EndBlock(store types.KVStore, height uint64) (res abci.ResponseEndBlock) {
	return
}
