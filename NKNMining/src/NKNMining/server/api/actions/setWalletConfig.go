package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"NKNMining/storage"
	"NKNMining/crypto"
	"fmt"
)

var SetWalletConfigAPI IRestfulAPIAction = &setWalletConfig{}

type setWalletConfigData struct {
	BeneficiaryAddr string
}

type setWalletConfig struct {
	restfulAPIBase
}

func (s *setWalletConfig) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/set/wallet/config"
}

func (s *setWalletConfig) Action(ctx *gin.Context) {
	response := apiServerResponse.New(ctx)

	inputJson, err := s.ExtractInput(ctx)
	if nil != err {
		response.BadRequest("invalid raw request!")
		return
	}

	basicData := &restfulAPIBaseRequest{}
	err = json.Unmarshal([]byte(inputJson), basicData)
	if nil != err {
		response.BadRequest("invalid request format!")
		return
	}

	walletInfoJsonStr, err := crypto.AesDecrypt(basicData.Data,
		storage.NKNSetupInfo.GetRequestKey())
	if nil != err {
		response.BadRequest("invalid request data!")
		return
	}

	walletInfoData := &setWalletConfigData{}
	err = json.Unmarshal([]byte(walletInfoJsonStr), walletInfoData)
	if nil != err {
		fmt.Println(err)
		response.BadRequest("invalid wallet setting data!")
		return
	}

	walletStorage := &storage.WalletConfig{}
	err = walletStorage.Save(walletInfoData.BeneficiaryAddr)
	if nil != err {
		// if checked {
		// 	response.InternalServerError("save wallet failed!")
		// } else {
		// 	response.BadRequest("invalid wallet data!")
		// }
		response.BadRequest(err)
	} else {
		
		response.Success(nil)
	}

	return
}
