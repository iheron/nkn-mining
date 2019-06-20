package storage

import (
	"NKNMining/common"
	"encoding/json"
	"io/ioutil"
)

const walletFile = "/bin/wallet.dat"

type Wallet struct {
	PasswordHash  string `json:"PasswordHash"`
	IV            string `json:"IV"`
	MasterKey     string `json:"MasterKey"`
	Version       string `json:"Version"`
	Address       string `json:"Address"`
	ProgramHash   string `json:"ProgramHash"`
	ContractData  string `json:"ContractData"`
	SeedEncrypted string `json:"SeedEncrypted"`
}

func (w *Wallet) Load() error {
	walletData, err := ioutil.ReadFile(common.GetCurrentDirectory() + walletFile)

	if nil != err {
		return err
	}

	json.Unmarshal(walletData, w)
	return nil
}

func (w *Wallet) simpleCheck() bool {
	return "" != w.PasswordHash &&
		"" != w.IV &&
		"" != w.MasterKey &&
		"" != w.Version &&
		"" != w.Address &&
		"" != w.ProgramHash &&
		"" != w.ContractData &&
		"" != w.SeedEncrypted
}

func (w *Wallet) Save(walletJsonStr string) (checked bool, err error) {
	checked = false
	newWallet := &Wallet{}
	err = json.Unmarshal([]byte(walletJsonStr), newWallet)
	if nil != err {
		return
	}

	if !newWallet.simpleCheck() {
		err = &common.NodeShellError{Code: common.NS_ERR_JSON_UNMARSHAL}
		return
	}

	checked = true
	err = ioutil.WriteFile(common.GetCurrentDirectory()+walletFile, []byte(walletJsonStr), 0666)
	return
}
