package apiServerAction

import (
	"NKNMining/common"
	"NKNMining/container"
	"NKNMining/server/api/const"
	"NKNMining/server/api/response"
	"github.com/gin-gonic/gin"
	"log"
)

var GetPublicKeyAPI IRestfulAPIAction = &getPublicKey{}

type getPublicKey struct {
	restfulAPIBase
}

func (g *getPublicKey) URI(serverURI string) string {
	return serverURI + apiServerConsts.API_SERVER_URI_BASE + "/get/publickey"
}

func (g *getPublicKey) Action(ctx *gin.Context) {
	response := apiServerResponse.New(ctx)
	publicKey := "..."

	publicKey, err := container.CmdApp.SyncRun([]string{"-c", "cat ./wallet.pswd | ./nknc wallet -l account |awk 'END{print $2}'"}, "")
	if nil != err {
		log.Println(err)
		publicKey = "UNKNOWN"
	}

	response.Success(map[string]string{
		"PublicKey":    publicKey,
		"ShellVersion": common.NS_VERSION,
		"GinVersion":   gin.Version,
	})

	return
}
