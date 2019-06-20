package storage

import (
	"NKNMining/common"
	"encoding/json"
	"io/ioutil"
)

const walletConfigFile = "/bin/config.json"

type WalletConfig struct {
	BeneficiaryAddr      string   `json:"BeneficiaryAddr,omitempty"`
	HttpWsPort           int      `json:"HttpWsPort"`
	HttpJsonPort         int      `json:"HttpJsonPort"`
	SeedList             []string `json:"SeedList"`
	GenesisBlockProposer string   `json:"GenesisBlockProposer"`
}

var WalletConfigInfo = &WalletConfig{}

func (w *WalletConfig) Load() error {
	walletData, err := ioutil.ReadFile(common.GetCurrentDirectory() + walletConfigFile)
	if nil != err {
		return err
	}
	// walletData = bytes.TrimPrefix(walletData, []byte("\xef\xbb\xbf"))
	json.Unmarshal(walletData, w)
	return nil
}

func (w *WalletConfig) Save(beneficiaryAddr string) error {
	WalletConfigInfo.Load()
	WalletConfigInfo.BeneficiaryAddr = beneficiaryAddr
	walletCfg, err := json.MarshalIndent(WalletConfigInfo, "", "    ")
	err = ioutil.WriteFile(common.GetCurrentDirectory()+walletConfigFile, walletCfg, 0666)
	return err
}
