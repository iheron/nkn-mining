package apiServerAction

import (
	"NKNMining/common"
	"NKNMining/container"
	"NKNMining/crypto"
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"NKNMining/storage"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var SetWalletAPI IRestfulAPIAction = &setWallet{}

type setWalletData struct {
	Wallet string
	Key    string
}

type setWallet struct {
	restfulAPIBase
}

func (s *setWallet) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/set/wallet"
}

func (s *setWallet) Action(ctx *gin.Context) {
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

	walletInfoData := &setWalletData{}
	err = json.Unmarshal([]byte(walletInfoJsonStr), walletInfoData)
	if nil != err {
		response.BadRequest("invalid wallet setting data!")
		return
	}

	walletStorage := &storage.Wallet{}
	checked, err := walletStorage.Save(walletInfoData.Wallet)
	if nil != err {
		if checked {
			response.InternalServerError("save wallet failed!")
		} else {
			response.BadRequest("invalid wallet data!")
		}
	} else {
		storage.NKNSetupInfo.CurrentStep = storage.SETUP_STEP_SUCCESS
		storage.NKNSetupInfo.WKey = walletInfoData.Key
		storage.NKNSetupInfo.Save()
		wKey, err := crypto.AesDecrypt(storage.NKNSetupInfo.WKey, storage.NKNSetupInfo.GetWalletKey())
		if nil != err {
			common.Log.Error(err)
			return
		}

		_, err = container.Node.AsyncRun([]string{"--no-check-port"}, wKey)
		//_, err = container.Node.SyncRun([]string{"--no-check-port"}, wKey)
		if nil != err {
			common.Log.Error(err)
			return
		}
		response.Success(nil)
	}

	return
}
