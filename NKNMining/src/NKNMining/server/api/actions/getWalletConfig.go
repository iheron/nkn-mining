package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	// "encoding/json"
	"NKNMining/storage"
)

var GetWalletConfigAPI IRestfulAPIAction = &getWalletConfig{}

type getWalletConfigData struct {
	BeneficiaryAddr string
}

type getWalletConfig struct {
	restfulAPIBase
}

func (s *getWalletConfig) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/get/wallet/config"
}

func (s *getWalletConfig) Action(ctx *gin.Context) {
	response := apiServerResponse.New(ctx)

	walletStorage := &storage.WalletConfig{}
	err := walletStorage.Load()
	if nil != err {
		// if checked {
		// 	response.InternalServerError("save wallet failed!")
		// } else {
		// 	response.BadRequest("invalid wallet data!")
		// }
		response.BadRequest(err)
	} else {
		
		response.Success(walletStorage)
	}

	return
}
