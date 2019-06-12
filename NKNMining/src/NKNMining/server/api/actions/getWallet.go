package apiServerAction

import (
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	// "encoding/json"
	"NKNMining/storage"
)

var GetWalletAPI IRestfulAPIAction = &getWallet{}

type getWalletData struct {
	BeneficiaryAddr string
}

type getWallet struct {
	restfulAPIBase
}

func (w *getWallet) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/get/wallet"
}

func (w *getWallet) Action(ctx *gin.Context) {
	response := apiServerResponse.New(ctx)

	walletStorage := &storage.Wallet{}
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
